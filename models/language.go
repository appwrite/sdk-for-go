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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Language) UnmarshalJSON(data []byte) error {
	type alias Language
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Name       json.RawMessage `json:"name"`
		Code       json.RawMessage `json:"code"`
		NativeName json.RawMessage `json:"nativeName"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Code == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Code); err != nil {
			return err
		}
	}
	if present.NativeName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.NativeName); err != nil {
			return err
		}
	}

	*model = Language(decoded)
	return nil
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
