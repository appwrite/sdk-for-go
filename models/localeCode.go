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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *LocaleCode) UnmarshalJSON(data []byte) error {
	type alias LocaleCode
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Code json.RawMessage `json:"code"`
		Name json.RawMessage `json:"name"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Code == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Code); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}

	*model = LocaleCode(decoded)
	return nil
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
