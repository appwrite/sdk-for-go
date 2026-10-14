package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasswordRotation Model
type PolicyPasswordRotation struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether expired passwords require a reset before email/password sign-in.
	Enabled bool `json:"enabled"`
	// Maximum password age in seconds.
	MaxAge int `json:"maxAge"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasswordRotation) New(data []byte) *PolicyPasswordRotation {
	model.data = data
	return &model
}

func (model *PolicyPasswordRotation) UnmarshalJSON(data []byte) error {
	type alias PolicyPasswordRotation
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false,\"maxAge\":31536000}"), &fields); err != nil {
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

func (model *PolicyPasswordRotation) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
