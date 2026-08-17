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
	// Overage and add-on pricing shared across all specifications.
	Pricing DedicatedDatabaseSpecificationPricing `json:"pricing"`

	// Used by Decode() method
	data []byte
}

func (model DedicatedDatabaseSpecificationList) New(data []byte) *DedicatedDatabaseSpecificationList {
	model.data = data
	return &model
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
