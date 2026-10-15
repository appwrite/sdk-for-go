package models

import (
	"encoding/json"
	"errors"
)

// Currency Model
type Currency struct {
	// Currency symbol.
	Symbol string `json:"symbol"`
	// Currency name.
	Name string `json:"name"`
	// Currency native symbol.
	SymbolNative string `json:"symbolNative"`
	// Number of decimal digits.
	DecimalDigits int `json:"decimalDigits"`
	// Currency digit rounding.
	Rounding float64 `json:"rounding"`
	// Currency code in [ISO 4217-1](http://en.wikipedia.org/wiki/ISO_4217)
	// three-character format.
	Code string `json:"code"`
	// Currency plural name
	NamePlural string `json:"namePlural"`

	// Used by Decode() method
	data []byte
}

func (model Currency) New(data []byte) *Currency {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Currency) UnmarshalJSON(data []byte) error {
	type alias Currency
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Symbol        json.RawMessage `json:"symbol"`
		Name          json.RawMessage `json:"name"`
		SymbolNative  json.RawMessage `json:"symbolNative"`
		DecimalDigits json.RawMessage `json:"decimalDigits"`
		Rounding      json.RawMessage `json:"rounding"`
		Code          json.RawMessage `json:"code"`
		NamePlural    json.RawMessage `json:"namePlural"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Symbol == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Symbol); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.SymbolNative == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SymbolNative); err != nil {
			return err
		}
	}
	if present.DecimalDigits == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.DecimalDigits); err != nil {
			return err
		}
	}
	if present.Rounding == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Rounding); err != nil {
			return err
		}
	}
	if present.Code == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Code); err != nil {
			return err
		}
	}
	if present.NamePlural == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.NamePlural); err != nil {
			return err
		}
	}

	*model = Currency(decoded)
	return nil
}

func (model *Currency) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
