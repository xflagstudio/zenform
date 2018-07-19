package config

import (
	"fmt"

	"github.com/nukosuke/go-zendesk/zendesk"
)

// map of slug and each resource instance
type TicketFieldStateMap map[string]zendesk.TicketField
type TicketFormStateMap map[string]zendesk.TicketForm
type TriggerStateMap map[string]zendesk.Trigger

// NotFoundID represents that the instance which has given slug does NOT exist
// When refering state map by slug, if the value not exists, map returns empty instance which ID is initialized by zero
const NotFoundID = 0

type ZenformState struct {
	TicketFields TicketFieldStateMap `json:"ticket_fields"`
	TicketForms  TicketFormStateMap  `json:"ticket_forms"`
	Triggers     TriggerStateMap     `json:"triggers"`
}

func NewZenformState() *ZenformState {
	return &ZenformState{
		TicketFields: TicketFieldStateMap{},
		TicketForms:  TicketFormStateMap{},
		Triggers:     TriggerStateMap{},
	}
}

// ActualConditionFieldName takes field name
// If given field name is system defined type, it returns argument itself
// If given field name is NOT system defined type, fetch custom field from ZenformState,
//   and make field name represented by `custom_fields_<id>`
func (state ZenformState) ActualConditionFieldName(systemFieldTypeOrSlug string) string {
	if zendesk.SystemConditionFieldTypes.Include(systemFieldTypeOrSlug) {
		return systemFieldTypeOrSlug
	}
	return fmt.Sprintf("custom_fields_%d", state.TicketFields[systemFieldTypeOrSlug].ID)
}

// ActualActionFieldName takes field name
// If given field name is system defined type, it returns argument itself
// If given field name is NOT system defined type, fetch custom field from ZenformState,
//   and make field name represented by `custom_fields_<id>`
func (state ZenformState) ActualActionFieldName(systemFieldTypeOrSlug string) string {
	if zendesk.SystemActionFieldTypes.Include(systemFieldTypeOrSlug) {
		return systemFieldTypeOrSlug
	}
	return fmt.Sprintf("custom_fields_%d", state.TicketFields[systemFieldTypeOrSlug].ID)
}

// ActualActionFieldValue takes field type and TriggerActionConfig.
// It returns field value according to field type
func (state ZenformState) ActualActionFieldValue(fieldType string, triggerAction TriggerActionConfig) interface{} {
	switch fieldType {
	case "notification_user": // []string
		return triggerAction.Value
	case "ticket_form_id": // convert slug to ticket form ID
		return state.TicketForms[triggerAction.Value[0]].ID
	default: // string
		return triggerAction.Value[0]
	}
}

// ExistsTicketField checks if given slug exists in the state map
func (state ZenformState) ExistsTicketField(slug string) bool {
	return state.TicketFields[slug].ID != NotFoundID
}

// ExistsTicketForm checks if given slug exists in the state map
func (state ZenformState) ExistsTicketForm(slug string) bool {
	return state.TicketForms[slug].ID != NotFoundID
}

// ExistsTrigger checks if given slug exists in the state map
func (state ZenformState) ExistsTrigger(slug string) bool {
	return state.Triggers[slug].ID != NotFoundID
}
