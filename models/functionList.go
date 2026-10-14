package models

import (
	"encoding/json"
	"errors"
)

// FunctionsList Model
type FunctionList struct {
	// Total number of functions that matched your query.
	Total int `json:"total"`
	// List of functions.
	Functions []Function `json:"functions"`

	// Used by Decode() method
	data []byte
}

func (model FunctionList) New(data []byte) *FunctionList {
	model.data = data
	return &model
}

func (model *FunctionList) UnmarshalJSON(data []byte) error {
	type alias FunctionList
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

func (model *FunctionList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
