package models

import (
	"encoding/json"
	"errors"
)

// ExecutionColumn Model
type DedicatedDatabaseExecutionColumn struct {
	// Column name as returned by the database.
	Name string `json:"name"`
	// Engine-specific column type (e.g. int4, text, timestamptz).
	Type string `json:"type"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseExecutionColumn) New(data []byte) *DedicatedDatabaseExecutionColumn {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabaseExecutionColumn) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseExecutionColumn
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Name json.RawMessage `json:"name"`
		Type json.RawMessage `json:"type"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}

	*model = DedicatedDatabaseExecutionColumn(decoded)
	return nil
}

func (model *DedicatedDatabaseExecutionColumn) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
