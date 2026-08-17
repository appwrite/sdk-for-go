package models

import (
	"encoding/json"
	"errors"
)

// OperationList Model
type DedicatedDatabaseOperationList struct {
	// Total number of operations.
	Total int `json:"total"`
	// List of operations.
	Operations []DedicatedDatabaseOperation `json:"operations"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseOperationList) New(data []byte) *DedicatedDatabaseOperationList {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseOperationList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
