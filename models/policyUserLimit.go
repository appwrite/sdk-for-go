package models

import (
	"encoding/json"
	"errors"
)

// PolicyUserLimit Model
type PolicyUserLimit struct {
	// Policy ID.
	Id string `json:"$id"`
	// Maximum number of users allowed in the project. A value of 0 means the
	// policy is disabled.
	Total int `json:"total"`

	// Used by Decode() method
	data []byte
}

func (model PolicyUserLimit) New(data []byte) *PolicyUserLimit {
	model.data = data
	return &model
}

func (model *PolicyUserLimit) UnmarshalJSON(data []byte) error {
	type alias PolicyUserLimit
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"total\":0}"), &fields); err != nil {
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

func (model *PolicyUserLimit) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
