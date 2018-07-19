package apply

import (
	"fmt"
	"strconv"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/xflagstudio/zenform/command/step"
	"github.com/xflagstudio/zenform/config"
)

func StepCreateTriggers(exe *step.Executor, zd *zendesk.Client, conf config.Config, state *config.ZenformState) func() error {
	return func() error {
		for _, trigger := range conf.Triggers {
			if state.ExistsTrigger(trigger.Slug) {
				continue
			}

			exe.Step("Creating \""+trigger.Title+"\"", func() error {
				//TODO: refactor, move to func
				position, err := strconv.ParseInt(trigger.Position, 10, 64)
				if err != nil {
					exe.Error(err.Error())
					return err
				}
				requestPayload := zendesk.Trigger{
					Title:    trigger.Title,
					Position: position,
				}

				for _, triggerAll := range trigger.All {
					allConditionPayload := zendesk.TriggerCondition{
						Field:    state.ActualConditionFieldName(triggerAll.Field),
						Operator: triggerAll.Operator,
						Value:    triggerAll.Value,
					}
					requestPayload.Conditions.All = append(requestPayload.Conditions.All, allConditionPayload)
				}

				for _, triggerAny := range trigger.Any {
					anyConditionPayload := zendesk.TriggerCondition{
						Field:    state.ActualConditionFieldName(triggerAny.Field),
						Operator: triggerAny.Operator,
						Value:    triggerAny.Value,
					}
					requestPayload.Conditions.Any = append(requestPayload.Conditions.Any, anyConditionPayload)
				}

				for _, triggerAction := range trigger.Actions {
					actionPayload := zendesk.TriggerAction{}
					actionPayload.Field = state.ActualActionFieldName(triggerAction.Field)
					actionPayload.Value = state.ActualActionFieldValue(actionPayload.Field, triggerAction)
					requestPayload.Actions = append(requestPayload.Actions, actionPayload)
				}

				result, err := zd.CreateTrigger(requestPayload)
				if err != nil {
					exe.Error("Failed to create trigger: " + trigger.Title)
					exe.Error("Error: " + err.Error())
					return err
				}
				state.Triggers[trigger.Slug] = result
				exe.Success(fmt.Sprintf("Done (ID=%d)", result.ID))
				return nil
			})
		}
		exe.Success("Created (" + fmt.Sprintf("%d", len(state.Triggers)) + ") triggers")
		return nil
	}
}
