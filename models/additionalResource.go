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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AdditionalResource) UnmarshalJSON(data []byte) error {
	type alias AdditionalResource
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Name        json.RawMessage `json:"name"`
		Unit        json.RawMessage `json:"unit"`
		Currency    json.RawMessage `json:"currency"`
		Price       json.RawMessage `json:"price"`
		Value       json.RawMessage `json:"value"`
		InvoiceDesc json.RawMessage `json:"invoiceDesc"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Unit == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Unit); err != nil {
			return err
		}
	}
	if present.Currency == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Currency); err != nil {
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

	*model = AdditionalResource(decoded)
	return nil
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
