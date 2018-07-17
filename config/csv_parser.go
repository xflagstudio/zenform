package config

import "encoding/json"

type CSVParser struct{}

func NewCSVParser() *CSVParser {
	return &CSVParser{}
}

func unmarshalJSONRow(jsonStr string) ([]string, error) {
	unmarshaledJSON := []string{}

	if jsonStr == "" {
		return unmarshaledJSON, nil
	}

	if err := json.Unmarshal([]byte(jsonStr), &unmarshaledJSON); err != nil {
		return []string{}, err
	}
	return unmarshaledJSON, nil
}

// isSameLength checks if fields, operators and values are array of same length
func isSameLength(conditionElements ...[]string) bool {
	expectedLength := len(conditionElements[0])
	for _, element := range conditionElements[1:] {
		if len(element) != expectedLength {
			return false
		}
	}
	return true
}
