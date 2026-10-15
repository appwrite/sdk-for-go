package models

import (
	"encoding/json"
	"errors"
)

// PostgresExtension Model
type PostgresExtension struct {
	// Extension key used with CREATE EXTENSION.
	Key string `json:"key"`
	// Human-readable extension name.
	Name string `json:"name"`
	// Short description of what the extension provides.
	Description string `json:"description"`
	// Category the extension belongs to.
	Category string `json:"category"`

	// Used by Decode() method
	data []byte
}

func (model PostgresExtension) New(data []byte) *PostgresExtension {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PostgresExtension) UnmarshalJSON(data []byte) error {
	type alias PostgresExtension
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Key         json.RawMessage `json:"key"`
		Name        json.RawMessage `json:"name"`
		Description json.RawMessage `json:"description"`
		Category    json.RawMessage `json:"category"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Key == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Key); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.Category == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Category); err != nil {
			return err
		}
	}

	*model = PostgresExtension(decoded)
	return nil
}

func (model *PostgresExtension) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
