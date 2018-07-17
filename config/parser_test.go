package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewParserFromExtension(t *testing.T) {
	assert := assert.New(t)
	supportedExtensions := []string{
		".csv",
	}

	for _, ext := range supportedExtensions {
		parser, _ := NewParserFromExtension(ext)
		assert.NotNil(parser)
	}

	parser, _ := NewParserFromExtension(".unknown_extension")
	assert.Nil(parser)
}
