package models

import (
	"encoding/json"
	"errors"
)

// Language Model
type Language struct {
	// Language name.
	Name string `json:"name"`
	// Language two-character ISO 639-1 codes.
	Code string `json:"code"`
	// Language native name.
	NativeName string `json:"nativeName"`

	// Used by Decode() method
	data []byte
}

func (model Language) New(data []byte) *Language {
	model.data = data
	return &model
}

func (model *Language) UnmarshalJSON(data []byte) error {
	type alias Language
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"name\":\"\",\"code\":\"\",\"nativeName\":\"\"}"), &fields); err != nil {
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

func (model *Language) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
