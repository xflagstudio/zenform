package apply

import (
	"fmt"
	"strconv"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/xflagstudio/zenform/command/step"
	"github.com/xflagstudio/zenform/config"
)

func StepCreateTicketForms(exe *step.Executor, zd *zendesk.Client, conf config.Config, state *config.ZenformState) func() error {
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
