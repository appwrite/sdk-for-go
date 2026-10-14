package models

import (
	"encoding/json"
	"errors"
)

// LocaleCode Model
type LocaleCode struct {
	// Locale codes in [ISO
	// 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes)
	Code string `json:"code"`
	// Locale name
	Name string `json:"name"`

	// Used by Decode() method
	data []byte
}

func (model LocaleCode) New(data []byte) *LocaleCode {
	model.data = data
	return &model
}

func (model *LocaleCode) UnmarshalJSON(data []byte) error {
	type alias LocaleCode
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"code\":\"\",\"name\":\"\"}"), &fields); err != nil {
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

func (model *LocaleCode) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
