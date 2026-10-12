package models

import (
	"encoding/json"
	"errors"
)

// FrameworksList Model
type FrameworkList struct {
	// Total number of frameworks that matched your query.
	Total int `json:"total"`
	// List of frameworks.
	Frameworks []Framework `json:"frameworks"`

	// Used by Decode() method
	data []byte
}

func (model FrameworkList) New(data []byte) *FrameworkList {
	model.data = data
	return &model
}

func (model *FrameworkList) UnmarshalJSON(data []byte) error {
	type alias FrameworkList
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

func (model *FrameworkList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
