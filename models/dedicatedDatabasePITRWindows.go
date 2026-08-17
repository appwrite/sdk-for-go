package models

import (
	"encoding/json"
	"errors"
)

// PITRWindows Model
type DedicatedDatabasePITRWindows struct {
	// Earliest available recovery point.
	Earliest string `json:"earliest"`
	// Latest available recovery point.
	Latest string `json:"latest"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabasePITRWindows) New(data []byte) *DedicatedDatabasePITRWindows {
	model.data = data
	return &model
}

func (model *DedicatedDatabasePITRWindows) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
