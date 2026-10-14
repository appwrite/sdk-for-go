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

func (model *Headers) UnmarshalJSON(data []byte) error {
	type alias Headers
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"name\":\"\",\"value\":\"\"}"), &fields); err != nil {
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
