package models

import (
	"encoding/json"
	"errors"
)

// Phone Model
type Phone struct {
	// Phone code.
	Code string `json:"code"`
	// Country two-character ISO 3166-1 alpha code.
	CountryCode string `json:"countryCode"`
	// Country name.
	CountryName string `json:"countryName"`

	// Used by Decode() method
	data []byte
}

func (model Phone) New(data []byte) *Phone {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Phone) UnmarshalJSON(data []byte) error {
	type alias Phone
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Code        json.RawMessage `json:"code"`
		CountryCode json.RawMessage `json:"countryCode"`
		CountryName json.RawMessage `json:"countryName"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Code == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Code); err != nil {
			return err
		}
	}
	if present.CountryCode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CountryCode); err != nil {
			return err
		}
	}
	if present.CountryName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CountryName); err != nil {
			return err
		}
	}

	*model = Phone(decoded)
	return nil
}

func (model *Phone) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
