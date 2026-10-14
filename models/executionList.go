package models

import (
	"encoding/json"
	"errors"
)

// ExecutionsList Model
type ExecutionList struct {
	// Total number of executions that matched your query.
	Total int `json:"total"`
	// List of executions.
	Executions []Execution `json:"executions"`

	// Used by Decode() method
	data []byte
}

func (model ExecutionList) New(data []byte) *ExecutionList {
	model.data = data
	return &model
}

func (model *ExecutionList) UnmarshalJSON(data []byte) error {
	type alias ExecutionList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
}

func (model *ExecutionList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
