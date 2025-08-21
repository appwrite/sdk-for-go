package models

import (
    "encoding/json"
    "errors"
)

// CurrenciesList Model
type CurrencyList struct {
    // Total number of currencies that matched your query.
    Total int `json:"total"`
    // List of currencies.
    Currencies []Currency `json:"currencies"`

    // Used by Decode() method
    data []byte
}

func (model CurrencyList) New(data []byte) *CurrencyList {
    model.data = data
    return &model
}

func (model *CurrencyList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}