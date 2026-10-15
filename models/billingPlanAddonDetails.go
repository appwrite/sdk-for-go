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
	Currency *string `json:"currency"`
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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *BillingPlanAddonDetails) UnmarshalJSON(data []byte) error {
	type alias BillingPlanAddonDetails
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Supported    json.RawMessage `json:"supported"`
		PlanIncluded json.RawMessage `json:"planIncluded"`
		Limit        json.RawMessage `json:"limit"`
		Type         json.RawMessage `json:"type"`
		Price        json.RawMessage `json:"price"`
		Value        json.RawMessage `json:"value"`
		InvoiceDesc  json.RawMessage `json:"invoiceDesc"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Supported == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Supported); err != nil {
			return err
		}
	}
	if present.PlanIncluded == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.PlanIncluded); err != nil {
			return err
		}
	}
	if present.Limit == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Limit); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"numeric\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Price == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Price); err != nil {
			return err
		}
	}
	if present.Value == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Value); err != nil {
			return err
		}
	}
	if present.InvoiceDesc == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.InvoiceDesc); err != nil {
			return err
		}
	}

	*model = BillingPlanAddonDetails(decoded)
	return nil
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
