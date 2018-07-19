package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/spf13/cobra"
	"github.com/xflagstudio/zenform/command/common"
	"github.com/xflagstudio/zenform/command/step"
	"github.com/xflagstudio/zenform/config"
)

var configFormat string

const stateFileName = "zfstate.json"

func init() {
	applyCommand.Flags().StringVarP(&configFormat, "format", "f", "csv", "Configuration file format. Only \"csv\" is supported currently.")
}

var applyCommand = &cobra.Command{
	Use:   "apply",
	Short: "Apply configuration to Zendesk",
	Long:  "Apply configuration to Zendesk",
	Run:   applyFunc,
}

func applyFunc(cmd *cobra.Command, args []string) {
	exe := step.NewExecutor()
	zd, _ := zendesk.NewClient(nil)
	zfconfig := &config.ZenformConfig{}
	currentState := config.NewZenformState()
	var conf config.Config

	// If specified zenform project directory path,
	// `apply` is executed in it.
	exe.Step("Running apply", func() error {
		if len(args) > 0 {
			os.Chdir(args[0])
		}
		cwd, _ := os.Getwd()
		exe.Info("Check into directory: " + cwd)
		return nil
	})

	exe.Step("Checking zenform config", common.StepLoadZenformConfig(exe, zfconfig, zd))

	// If zfstate.json exists,
	// recover last config data as current Zendesk state
	exe.Step("Checking state file", func() error {
		if _, err := os.Stat(stateFileName); os.IsNotExist(err) {
			exe.Info("Not Found")
			return nil
		}
		exe.Success("Found " + stateFileName)

		exe.Step("Loading state", func() error {
			stateJsonStr, _ := ioutil.ReadFile(stateFileName)
			err := json.Unmarshal(stateJsonStr, &currentState)
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			return nil
		})
		return nil
	})

	// Create parser according to config extension
	configExtension := "." + configFormat
	parser, err := config.NewParserFromExtension(configExtension)
	if err != nil {
		fmt.Println(err)
		return
	}

	//TODO: check if all configs exists
	exe.Step("Detecting config files", func() error {
		exe.Success("Found (number of files) files")
		return nil
	})

	exe.Step("Loading configuration files", func() error {
		exe.Step("Loading ticket_fields"+configExtension, func() error {
			confText, err := ioutil.ReadFile("./ticket_fields" + configExtension)
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			conf.TicketFields, err = parser.ParseTicketFields(string(confText))
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			exe.Success("OK")
			return nil
		})

		exe.Step("Loading ticket_forms"+configExtension, func() error {
			confText, err := ioutil.ReadFile("./ticket_forms" + configExtension)
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			conf.TicketForms, err = parser.ParseTicketForms(string(confText))
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			exe.Success("OK")
			return nil
		})

		exe.Step("Loading triggers"+configExtension, func() error {
			confText, err := ioutil.ReadFile("./triggers" + configExtension)
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			conf.Triggers, err = parser.ParseTriggers(string(confText))
			if err != nil {
				exe.Error(err.Error())
				return err
			}
			exe.Success("OK")
			return nil
		})
		return nil
	})

	exe.Step("Applying patch to Zendesk", func() error {
		if err := exe.StepIf(len(conf.TicketFields) > 0, "Creating ticket fields...", stepCreateTicketFields(exe, zd, conf, currentState)); err != nil {
			return err
		}
		if err := exe.StepIf(len(conf.TicketForms) > 0, "Creating ticket forms...", stepCreateTicketForms(exe, zd, conf, currentState)); err != nil {
			return err
		}
		if err := exe.StepIf(len(conf.Triggers) > 0, "Creating triggers...", stepCreateTriggers(exe, zd, conf, currentState)); err != nil {
			return err
		}
		return nil
	})

	// Write new state file
	exe.Step("Updating zfstate.json", func() error {
		jsonStr, err := json.MarshalIndent(currentState, "", "    ") // indent with 4 spaces
		if err != nil {
			exe.Error(err.Error())
			return err
		}
		ioutil.WriteFile(stateFileName, jsonStr, 0644)
		exe.Success("Done")
		return nil
	})
}

func stepCreateTicketFields(exe *step.Executor, zd *zendesk.Client, conf config.Config, state *config.ZenformState) func() error {
	return func() error {
		for _, ticketField := range conf.TicketFields {
			if state.ExistsTicketField(ticketField.Slug) {
				continue
			}

			exe.Step("Creating \""+ticketField.Title+"\"", func() error {
				//TODO: refactor, move to func
				requestPayload := zendesk.TicketField{
					Type:  ticketField.Type,
					Title: ticketField.Title,
				}
				for _, customFieldOption := range ticketField.CustomFieldOptions {
					requestPayload.CustomFieldOptions = append(requestPayload.CustomFieldOptions, zendesk.TicketFieldCustomFieldOption{
						Name:  customFieldOption.Name,
						Value: customFieldOption.Value,
					})
				}

				result, err := zd.CreateTicketField(requestPayload)
				if err != nil {
					exe.Error("Failed to create ticket field: " + ticketField.Title)
					exe.Error("Error: " + err.Error())
					return err
				}
				state.TicketFields[ticketField.Slug] = result
				exe.Success(fmt.Sprintf("Done (ID=%d)", result.ID))
				return nil
			})
		}
		exe.Success("Created (" + fmt.Sprintf("%d", len(state.TicketFields)) + ") ticket fields")
		return nil
	}
}

func stepCreateTicketForms(exe *step.Executor, zd *zendesk.Client, conf config.Config, state *config.ZenformState) func() error {
	return func() error {
		for _, ticketForm := range conf.TicketForms {
			if state.ExistsTicketForm(ticketForm.Slug) {
				continue
			}

			exe.Step("Creating \""+ticketForm.Name+"\"", func() error {
				//TODO: refactor, move to func
				position, err := strconv.ParseInt(ticketForm.Position, 10, 64)
				if err != nil {
					exe.Error(err.Error())
					return err
				}
				// get actual ticket field ID from ZenformState by slug
				ticketFieldIDs := []int64{}
				for _, ticketFieldIDString := range ticketForm.TicketFieldIDs {
					ticketFieldIDs = append(ticketFieldIDs, state.TicketFields[ticketFieldIDString].ID)
				}
				requestPayload := zendesk.TicketForm{
					Name:           ticketForm.Name,
					Position:       position,
					TicketFieldIDs: ticketFieldIDs,
				}

				result, err := zd.CreateTicketForm(requestPayload)
				if err != nil {
					exe.Error("Failed to create ticket form: " + ticketForm.Name)
					exe.Error("Error: " + err.Error())
					return err
				}
				state.TicketForms[ticketForm.Slug] = result
				exe.Success(fmt.Sprintf("Done (ID=%d)", result.ID))
				return nil
			})
		}
		exe.Success("Created (" + fmt.Sprintf("%d", len(state.TicketForms)) + ") ticket forms")
		return nil
	}
}

func stepCreateTriggers(exe *step.Executor, zd *zendesk.Client, conf config.Config, state *config.ZenformState) func() error {
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
