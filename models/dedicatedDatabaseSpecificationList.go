package models

import (
	"encoding/json"
	"errors"
)

// SpecificationList Model
type DedicatedDatabaseSpecificationList struct {
	// List of dedicated database specifications.
	Specifications []DedicatedDatabaseSpecification `json:"specifications"`
	// Total number of specifications.
	Total int `json:"total"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseSpecificationList) New(data []byte) *DedicatedDatabaseSpecificationList {
	model.data = data
	return &model
}

func (model *DedicatedDatabaseSpecificationList) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabaseSpecificationList
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

func (model *DedicatedDatabaseSpecificationList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
