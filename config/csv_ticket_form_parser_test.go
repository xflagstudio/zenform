package config

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const ticketFormsFixture = "../test/fixtures/ticket_forms.csv"

func TestParseTicketForms(t *testing.T) {
	assert := assert.New(t)
	parser, _ := NewParserFromExtension(".csv")
	csvStr, _ := ioutil.ReadFile(ticketFormsFixture)
	conf, err := parser.ParseTicketForms(string(csvStr))
	expectedLength := 7 // number of rows except header in ticket_forms.csv

	assert.Nil(err)
	assert.Equal(len(conf), expectedLength)
}
