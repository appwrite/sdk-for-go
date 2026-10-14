package models

import (
	"encoding/json"
	"errors"
)

// PolicyDenyFreeEmail Model
type PolicyDenyFreeEmail struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether the deny free email policy is enabled.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model PolicyDenyFreeEmail) New(data []byte) *PolicyDenyFreeEmail {
	model.data = data
	return &model
}

func (model *PolicyDenyFreeEmail) UnmarshalJSON(data []byte) error {
	type alias PolicyDenyFreeEmail
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false}"), &fields); err != nil {
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

func (model *PolicyDenyFreeEmail) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
