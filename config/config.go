package config

type Config struct {
	TicketFields []TicketFieldConfig
	TicketForms  []TicketFormConfig
	Triggers     []TriggerConfig
}

type TicketFieldConfig struct {
	Slug               string
	Title              string
	Type               string
	VisibleInPortal    string
	EditableInPortal   string
	RequiredInPortal   string
	Description        string
	CustomFieldOptions []TicketFieldCustomFieldOptionConfig
}

type TicketFieldCustomFieldOptionConfig struct {
	Name  string
	Value string
}

type TicketFormConfig struct {
	Slug           string
	Name           string
	Position       string
	EndUserVisible string
	DisplayName    string
	TicketFieldIDs []string
}

type TriggerConfig struct {
	Slug     string
	Title    string
	Position string
	All      []TriggerConditionConfig
	Any      []TriggerConditionConfig
	Actions  []TriggerActionConfig
}

type TriggerConditionConfig struct {
	Field    string
	Operator string
	Value    string
}

type TriggerActionConfig struct {
	Field string
	Value []string
}
