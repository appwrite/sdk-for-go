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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AttributeList) UnmarshalJSON(data []byte) error {
	type alias AttributeList
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Total json.RawMessage `json:"total"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Total == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Total); err != nil {
			return err
		}
	}

	*model = AttributeList(decoded)
	return nil
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
