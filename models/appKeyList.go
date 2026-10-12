package models

import (
	"encoding/json"
	"errors"
)

// AppKeysList Model
type AppKeyList struct {
	// Total number of keys that matched your query.
	Total int `json:"total"`
	// List of keys.
	Keys []AppKey `json:"keys"`

	// Used by Decode() method
	data []byte
}

func (model AppKeyList) New(data []byte) *AppKeyList {
	model.data = data
	return &model
}

func (model *AppKeyList) UnmarshalJSON(data []byte) error {
	type alias AppKeyList
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

func (model *AppKeyList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
