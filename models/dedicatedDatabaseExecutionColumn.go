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
