package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const ticketFieldsFixture = "../test/fixtures/ticket_fields.csv"

func TestParseTicketFields(t *testing.T) {
	assert := assert.New(t)
	parser, _ := NewParserFromExtension(".csv")
	csvStr, _ := ioutil.ReadFile(ticketFieldsFixture)
	conf, err := parser.ParseTicketFields(string(csvStr))
	expectedLength := 28 // number of rows except header in ticket_fields.csv
	fmt.Println(conf)

	assert.Nil(err)
	assert.Equal(len(conf), expectedLength)
}
