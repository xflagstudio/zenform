package config

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type ticketFieldColumn int

const (
	colTicketFieldSlug ticketFieldColumn = iota
	colTicketFieldTitle
	colTicketFieldType
	colTicketFieldVisibleInPortal
	colTicketFieldEditableInPortal
	colTicketFieldRequiredInPortal
	colTicketFieldDescription
	colTicketFieldCustomFieldOptionsNames
	colTicketFieldCustomFieldOptionsValues
)

func (p CSVParser) ParseTicketFields(csvStr string) ([]TicketFieldConfig, error) {
	ticketFields := []TicketFieldConfig{}
	reader := csv.NewReader(strings.NewReader(csvStr))
	reader.Comma = ','

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for index, row := range rows[1:] {
		linum := index + 2 // start from 1 and except header => 2
		ticketField, err := transformToTicketField(row)

		if err != nil {
			return nil, fmt.Errorf("line %d: %s", linum, err)
		}
		ticketFields = append(ticketFields, ticketField)
	}

	return ticketFields, nil
}

func transformToTicketField(row []string) (TicketFieldConfig, error) {
	ticketField := TicketFieldConfig{
		Slug:             row[colTicketFieldSlug],
		Title:            row[colTicketFieldTitle],
		Type:             row[colTicketFieldType],
		VisibleInPortal:  row[colTicketFieldVisibleInPortal],
		EditableInPortal: row[colTicketFieldEditableInPortal],
		RequiredInPortal: row[colTicketFieldRequiredInPortal],
		Description:      row[colTicketFieldDescription],
	}

	customFieldOptions, err := transformToTicketFieldCustomFieldOption(row[colTicketFieldCustomFieldOptionsNames], row[colTicketFieldCustomFieldOptionsValues])
	if err != nil {
		return TicketFieldConfig{}, err
	}
	ticketField.CustomFieldOptions = customFieldOptions

	return ticketField, nil
}

func transformToTicketFieldCustomFieldOption(namesJSON string, valuesJSON string) ([]TicketFieldCustomFieldOptionConfig, error) {
	fieldOptions := []TicketFieldCustomFieldOptionConfig{}

	names, err := unmarshalJSONRow(namesJSON)
	if err != nil {
		return fieldOptions, err
	}

	values, err := unmarshalJSONRow(valuesJSON)
	if err != nil {
		return fieldOptions, err
	}

	fieldOption := TicketFieldCustomFieldOptionConfig{}
	iterationCount := len(names)
	for i := 0; i < iterationCount; i++ {
		fieldOption.Name = names[i]
		fieldOption.Value = values[i]
		fieldOptions = append(fieldOptions, fieldOption)
	}
	return fieldOptions, nil
}
