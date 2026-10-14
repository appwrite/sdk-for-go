package models

import (
	"encoding/json"
	"errors"
)

// AppSecretsList Model
type AppSecretList struct {
	// Total number of secrets that matched your query.
	Total int `json:"total"`
	// List of secrets.
	Secrets []AppSecret `json:"secrets"`

	// Used by Decode() method
	data []byte
}

func (model AppSecretList) New(data []byte) *AppSecretList {
	model.data = data
	return &model
}

func (model *AppSecretList) UnmarshalJSON(data []byte) error {
	type alias AppSecretList
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

func (model *AppSecretList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
