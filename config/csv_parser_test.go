package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalJSONRow(t *testing.T) {
	assert := assert.New(t)

	jsonStr := "[\"value1\", \"value2\"]"
	result, _ := unmarshalJSONRow(jsonStr)
	assert.Equal(result, []string{"value1", "value2"})

	emptyStr := ""
	result, _ = unmarshalJSONRow(emptyStr)
	assert.Equal(result, []string{})

	invalidStr := "[\"value1\", \"value2\"" // missing closing bracket
	_, err := unmarshalJSONRow(invalidStr)
	assert.NotNil(err)
}

func TestIsSameLength(t *testing.T) {
	assert := assert.New(t)

	elements := [][]string{
		[]string{"key1", "key2"},
		[]string{"value1", "value2"},
	}
	assert.True(isSameLength(elements...))

	elements = [][]string{
		[]string{"key1", "key2", "key3"},
		[]string{"value1", "value2"},
	}
	assert.False(isSameLength(elements...))
}
