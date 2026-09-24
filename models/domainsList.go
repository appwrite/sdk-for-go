package models

import (
	"encoding/json"
	"errors"
)

// DomainsList Model
type DomainsList struct {
	// Total number of domains that matched your query.
	Total int `json:"total"`
	// List of domains.
	Domains []Domain `json:"domains"`

	// Used by Decode() method
	data []byte
}

func (model DomainsList) New(data []byte) *DomainsList {
	model.data = data
	return &model
}

func (model *DomainsList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
