package models

import (
	"encoding/json"
	"errors"
)

// Continent Model
type Continent struct {
	// Continent name.
	Name string `json:"name"`
	// Continent two letter code.
	Code string `json:"code"`

	// Used by Decode() method
	data []byte
}

func (model Continent) New(data []byte) *Continent {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Continent) UnmarshalJSON(data []byte) error {
	type alias Continent
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Name json.RawMessage `json:"name"`
		Code json.RawMessage `json:"code"`
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

	*model = Continent(decoded)
	return nil
}

func (model *Continent) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
