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