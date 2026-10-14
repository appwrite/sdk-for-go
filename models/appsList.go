package models

import (
	"encoding/json"
	"errors"
)

// AppsList Model
type AppsList struct {
	// Total number of apps that matched your query.
	Total int `json:"total"`
	// List of apps.
	Apps []App `json:"apps"`

	// Used by Decode() method
	data []byte
}

func (model AppsList) New(data []byte) *AppsList {
	model.data = data
	return &model
}

func (model *AppsList) UnmarshalJSON(data []byte) error {
	type alias AppsList
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

func (model *AppsList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
