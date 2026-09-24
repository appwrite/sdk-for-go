package models

import (
	"encoding/json"
	"errors"
)

// DomainPrice Model
type DomainPrice struct {
	// Domain name.
	Domain string `json:"domain"`
	// Top-level domain for the requested domain.
	Tld string `json:"tld"`
	// Whether the domain is currently available for registration.
	Available bool `json:"available"`
	// Domain registration price. Null when the price could not be resolved, for
	// example for an unsupported TLD.
	Price float64 `json:"price"`
	// Price period in years.
	PeriodYears int `json:"periodYears"`
	// Whether the domain is a premium domain.
	Premium bool `json:"premium"`
	// Domain renewal price for the same period. Null when the domain was not
	// priced or the registrar has no renewal price for it.
	RenewalPrice float64 `json:"renewalPrice"`
	// Renewal price period in years.
	RenewalPeriodYears int `json:"renewalPeriodYears"`

	// Used by Decode() method
	data []byte
}

func (model DomainPrice) New(data []byte) *DomainPrice {
	model.data = data
	return &model
}

func (model *DomainPrice) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
