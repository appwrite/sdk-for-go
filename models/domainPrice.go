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
	Price *float64 `json:"price"`
	// Price period in years.
	PeriodYears int `json:"periodYears"`
	// Whether the domain is a premium domain.
	Premium bool `json:"premium"`
	// Domain renewal price for the same period. Null when the domain was not
	// priced or the registrar has no renewal price for it.
	RenewalPrice *float64 `json:"renewalPrice"`
	// Renewal price period in years.
	RenewalPeriodYears int `json:"renewalPeriodYears"`

	// Used by Decode() method
	data []byte
}

func (model DomainPrice) New(data []byte) *DomainPrice {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DomainPrice) UnmarshalJSON(data []byte) error {
	type alias DomainPrice
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Domain             json.RawMessage `json:"domain"`
		Tld                json.RawMessage `json:"tld"`
		Available          json.RawMessage `json:"available"`
		PeriodYears        json.RawMessage `json:"periodYears"`
		Premium            json.RawMessage `json:"premium"`
		RenewalPeriodYears json.RawMessage `json:"renewalPeriodYears"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Domain == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Domain); err != nil {
			return err
		}
	}
	if present.Tld == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Tld); err != nil {
			return err
		}
	}
	if present.Available == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Available); err != nil {
			return err
		}
	}
	if present.PeriodYears == nil {
		if err := json.Unmarshal([]byte("1"), &decoded.PeriodYears); err != nil {
			return err
		}
	}
	if present.Premium == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Premium); err != nil {
			return err
		}
	}
	if present.RenewalPeriodYears == nil {
		if err := json.Unmarshal([]byte("1"), &decoded.RenewalPeriodYears); err != nil {
			return err
		}
	}

	*model = DomainPrice(decoded)
	return nil
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
