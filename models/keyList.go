package models

import (
	"encoding/json"
	"errors"
)

// APIKeysList Model
type KeyList struct {
	// Total number of keys that matched your query.
	Total int `json:"total"`
	// List of keys.
	Keys []Key `json:"keys"`

	// Used by Decode() method
	data []byte
}

func (model KeyList) New(data []byte) *KeyList {
	model.data = data
	return &model
}

func (model *KeyList) UnmarshalJSON(data []byte) error {
	type alias KeyList
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

func (model *KeyList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
