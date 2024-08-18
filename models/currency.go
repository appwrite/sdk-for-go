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