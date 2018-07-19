package apply

import (
	"fmt"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/xflagstudio/zenform/command/step"
	"github.com/xflagstudio/zenform/config"
)

func StepCreateTicketFields(exe *step.Executor, zd *zendesk.Client, conf config.Config, state *config.ZenformState) func() error {
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
