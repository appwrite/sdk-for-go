package models

import (
	"encoding/json"
	"errors"
)

// Extensions Model
type DedicatedDatabaseExtensions struct {
	// List of installed extensions.
	Installed []string `json:"installed"`
	// List of available extensions that can be installed.
	Available []string `json:"available"`
	// Curated metadata (display name, description, category) for each available
	// extension.
	Metadata []PostgresExtension `json:"metadata"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseExtensions) New(data []byte) *DedicatedDatabaseExtensions {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseExtensions) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
