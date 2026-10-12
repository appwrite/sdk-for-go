package models

import (
	"encoding/json"
	"errors"
)

// AttributesList Model
type AttributeList struct {
	// Total number of attributes in the given collection.
	Total int `json:"total"`
	// List of attributes.
	Attributes []map[string]any `json:"attributes"`

	// Used by Decode() method
	data []byte
}

func (model AttributeList) New(data []byte) *AttributeList {
	model.data = data
	return &model
}

func (model *AttributeList) UnmarshalJSON(data []byte) error {
	type alias AttributeList
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

func (model *AttributeList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
