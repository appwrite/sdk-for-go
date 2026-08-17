package models

import (
	"encoding/json"
	"errors"
)

// BranchList Model
type DedicatedDatabaseBranchList struct {
	// List of branches.
	Branches []DedicatedDatabaseBranch `json:"branches"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseBranchList) New(data []byte) *DedicatedDatabaseBranchList {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseBranchList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
