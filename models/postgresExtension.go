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

func (model *PostgresExtension) UnmarshalJSON(data []byte) error {
	type alias PostgresExtension
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"key\":\"\",\"name\":\"\",\"description\":\"\",\"category\":\"\"}"), &fields); err != nil {
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
