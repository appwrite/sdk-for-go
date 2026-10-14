package models

import (
	"encoding/json"
	"errors"
)

// ProjectAuthMethod Model
type ProjectAuthMethod struct {
	// Auth method ID.
	Id string `json:"$id"`
	// Auth method status.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model ProjectAuthMethod) New(data []byte) *ProjectAuthMethod {
	model.data = data
	return &model
}

func (model *ProjectAuthMethod) UnmarshalJSON(data []byte) error {
	type alias ProjectAuthMethod
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"enabled\":true}"), &fields); err != nil {
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

func (model *ProjectAuthMethod) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
