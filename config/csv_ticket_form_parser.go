package config

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type ticketFormColumn int

const (
	colTicketFormSlug ticketFormColumn = iota
	colTicketFormName
	colTicketFormPosition
	colTicketFormEndUserVisible
	colTicketFormDisplayName
	colTicketFormTicketFieldIDs
)

func (p CSVParser) ParseTicketForms(csvStr string) ([]TicketFormConfig, error) {
	ticketForms := []TicketFormConfig{}
	reader := csv.NewReader(strings.NewReader(csvStr))
	reader.Comma = ','

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for index, row := range rows[1:] {
		linum := index + 2 // start from 1 and except header => 2
		ticketForm, err := transformToTicketForm(row)

		if err != nil {
			return nil, fmt.Errorf("line %d: %s", linum, err)
		}
		ticketForms = append(ticketForms, ticketForm)
	}
	return ticketForms, nil
}

func transformToTicketForm(row []string) (TicketFormConfig, error) {
	ticketForm := TicketFormConfig{
		Slug:           row[colTicketFormSlug],
		Name:           row[colTicketFormName],
		Position:       row[colTicketFormPosition],
		EndUserVisible: row[colTicketFormEndUserVisible],
		DisplayName:    row[colTicketFormDisplayName],
	}

	ticketFieldIDs, err := unmarshalJSONRow(row[colTicketFormTicketFieldIDs])
	if err != nil {
		return TicketFormConfig{}, err
	}
	ticketForm.TicketFieldIDs = ticketFieldIDs

	return ticketForm, nil
}
