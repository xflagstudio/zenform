package config

import "errors"

type Parser interface {
	ParseTicketFields(string) ([]TicketFieldConfig, error)
	ParseTicketForms(string) ([]TicketFormConfig, error)
	ParseTriggers(string) ([]TriggerConfig, error)
}

func NewParserFromExtension(ext string) (Parser, error) {
	switch ext {
	case ".csv":
		return NewCSVParser(), nil
	default:
		return nil, errors.New("Given extension '" + ext + "' is not supported")
	}
}
