package models

import (
	"encoding/json"
	"errors"
)

// PoliciesList Model
type PolicyList struct {
	// Total number of policies in the given project.
	Total int `json:"total"`
	// List of policies.
	Policies []interface{} `json:"policies"`

	// Used by Decode() method
	data []byte
}

func (model PolicyList) New(data []byte) *PolicyList {
	model.data = data
	return &model
}

func (model *PolicyList) UnmarshalJSON(data []byte) error {
	type alias PolicyList
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

func (model *PolicyList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
