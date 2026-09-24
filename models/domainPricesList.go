package models

import (
	"encoding/json"
	"errors"
)

// DomainPricesList Model
type DomainPricesList struct {
	// Total number of prices that matched your query.
	Total int `json:"total"`
	// List of prices.
	Prices []DomainPrice `json:"prices"`

	// Used by Decode() method
	data []byte
}

func (model DomainPricesList) New(data []byte) *DomainPricesList {
	model.data = data
	return &model
}

func (model *DomainPricesList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
