package models

import (
	"encoding/json"
	"errors"
)

// BranchList Model
type DedicatedDatabaseBranchList struct {
	// Total number of branches.
	Total int `json:"total"`
	// List of branches.
	Branches []DedicatedDatabaseBranch `json:"branches"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseBranchList) New(data []byte) *DedicatedDatabaseBranchList {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseBranchList) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseBranchList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
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
