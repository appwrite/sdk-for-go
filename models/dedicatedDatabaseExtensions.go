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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabaseExtensions) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseExtensions
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Installed json.RawMessage `json:"installed"`
		Available json.RawMessage `json:"available"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Installed == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Installed); err != nil {
			return err
		}
	}
	if present.Available == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Available); err != nil {
			return err
		}
	}

	*model = DedicatedDatabaseExtensions(decoded)
	return nil
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
