package config

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

type triggerColumn int

const (
	colTriggerSlug triggerColumn = iota
	colTriggerTitle
	colTriggerPosition
	colTriggerConditionsAllFields
	colTriggerConditionsAllOperators
	colTriggerConditionsAllValues
	colTriggerConditionsAnyFields
	colTriggerConditionsAnyOperators
	colTriggerConditionsAnyValues
	colTriggerActionsFields
	colTriggerActionsValues
)

func (p CSVParser) ParseTriggers(csvStr string) ([]TriggerConfig, error) {
	triggers := []TriggerConfig{}
	reader := csv.NewReader(strings.NewReader(csvStr))
	reader.Comma = ','

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for index, row := range rows[1:] {
		linum := index + 2 // start from 1 and except header => 2
		trigger, err := transformToTrigger(row)

		if err != nil {
			return nil, fmt.Errorf("line %d: %s", linum, err)
		}
		triggers = append(triggers, trigger)
	}
	return triggers, nil
}

func transformToTrigger(row []string) (TriggerConfig, error) {
	trigger := TriggerConfig{
		Slug:     row[colTriggerSlug],
		Title:    row[colTriggerTitle],
		Position: row[colTriggerPosition],
	}

	allConditions, err := transformToTriggerCondition(
		row[colTriggerConditionsAllFields],
		row[colTriggerConditionsAllOperators],
		row[colTriggerConditionsAllValues],
	)

	if err != nil {
		return TriggerConfig{}, err
	}
	trigger.All = allConditions

	anyConditions, err := transformToTriggerCondition(
		row[colTriggerConditionsAnyFields],
		row[colTriggerConditionsAnyOperators],
		row[colTriggerConditionsAnyValues],
	)

	if err != nil {
		return TriggerConfig{}, err
	}
	trigger.Any = anyConditions

	actions, err := transformToTriggerAction(row[colTriggerActionsFields], row[colTriggerActionsValues])
	if err != nil {
		return TriggerConfig{}, err
	}
	trigger.Actions = actions

	return trigger, nil
}

// transformToTriggerCondition takes columns include fields, operators and values in JSON string
// and returns condition which is transforrmed into array of TriggerConditionConfig struct
func transformToTriggerCondition(fieldsJSON string, operatorsJSON string, valuesJSON string) ([]TriggerConditionConfig, error) {
	conditions := []TriggerConditionConfig{}

	fields, err := unmarshalJSONRow(fieldsJSON)
	if err != nil {
		return conditions, err
	}

	operators, err := unmarshalJSONRow(operatorsJSON)
	if err != nil {
		return conditions, err
	}

	values, err := unmarshalJSONRow(valuesJSON)
	if err != nil {
		return conditions, err
	}

	if !isSameLength(fields, operators, values) {
		return conditions, fmt.Errorf("number of condition's element does not match")
	}

	// convert field, operator and value arrays into structured data
	condition := TriggerConditionConfig{}
	iterationCount := len(fields)
	for i := 0; i < iterationCount; i++ {
		condition.Field = fields[i]
		condition.Operator = operators[i]
		condition.Value = values[i]
		conditions = append(conditions, condition)
	}
	return conditions, nil
}

func transformToTriggerAction(fieldsJSON string, valuesJSON string) ([]TriggerActionConfig, error) {
	actions := []TriggerActionConfig{}

	fields, err := unmarshalJSONRow(fieldsJSON)
	if err != nil {
		return actions, err
	}

	values := [][]string{}
	if err := json.Unmarshal([]byte(valuesJSON), &values); err != nil {
		return actions, err
	}

	if len(fields) != len(values) {
		return actions, fmt.Errorf("number of action's fields and values does not match")
	}

	action := TriggerActionConfig{}
	iterationCount := len(fields)
	for i := 0; i < iterationCount; i++ {
		action.Field = fields[i]
		action.Value = values[i]
		actions = append(actions, action)
	}
	return actions, nil
}
