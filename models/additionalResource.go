package models

import (
	"encoding/json"
	"errors"
)

// AdditionalResource Model
type AdditionalResource struct {
	// Resource name
	Name string `json:"name"`
	// Resource unit
	Unit string `json:"unit"`
	// Price currency
	Currency string `json:"currency"`
	// Price
	Price float64 `json:"price"`
	// Resource value
	Value int `json:"value"`
	// Description on invoice
	InvoiceDesc string `json:"invoiceDesc"`

	// Used by Decode() method
	data []byte
}

func (model AdditionalResource) New(data []byte) *AdditionalResource {
	model.data = data
	return &model
}

func (model *AdditionalResource) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
