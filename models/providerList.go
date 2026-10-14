package models

import (
	"encoding/json"
	"errors"
)

// ProviderList Model
type ProviderList struct {
	// Total number of providers that matched your query.
	Total int `json:"total"`
	// List of providers.
	Providers []Provider `json:"providers"`

	// Used by Decode() method
	data []byte
}

func (model ProviderList) New(data []byte) *ProviderList {
	model.data = data
	return &model
}

func (model *ProviderList) UnmarshalJSON(data []byte) error {
	type alias ProviderList
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

func (model *ProviderList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
