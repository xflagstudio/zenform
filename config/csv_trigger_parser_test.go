package config

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const fixture = "../test/fixtures/triggers.csv"

func TestParseTriggers(t *testing.T) {
	assert := assert.New(t)
	parser, _ := NewParserFromExtension(".csv")
	csvStr, _ := ioutil.ReadFile(fixture)
	conf, err := parser.ParseTriggers(string(csvStr))
	expectedLength := 5 // number of rows except header in triggers.csv

	assert.Nil(err)
	assert.Equal(len(conf), expectedLength)
}
