package models

import (
	"encoding/json"
	"errors"
)

// Details Model
type BillingPlanAddonDetails struct {
	// Is the addon supported in the plan?
	Supported bool `json:"supported"`
	// Addon plan included value
	PlanIncluded int `json:"planIncluded"`
	// Addon limit
	Limit int `json:"limit"`
	// Addon type
	Type string `json:"type"`
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

func (model BillingPlanAddonDetails) New(data []byte) *BillingPlanAddonDetails {
	model.data = data
	return &model
}

func (model *BillingPlanAddonDetails) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
