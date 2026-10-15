package models

import (
	"encoding/json"
	"errors"
)

// Headers Model
type Headers struct {
	// Header name.
	Name string `json:"name"`
	// Header value.
	Value string `json:"value"`

	// Used by Decode() method
	data []byte
}

func (model Headers) New(data []byte) *Headers {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Headers) UnmarshalJSON(data []byte) error {
	type alias Headers
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Name  json.RawMessage `json:"name"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Value == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Value); err != nil {
			return err
		}
	}

	*model = Headers(decoded)
	return nil
}

func (model *Headers) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
