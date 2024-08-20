package models

import (
    "encoding/json"
    "errors"
)

// Continent Model
type Continent struct {
    // Continent name.
    Name string `json:"name"`
    // Continent two letter code.
    Code string `json:"code"`

    // Used by Decode() method
    data []byte
}

func (model Continent) New(data []byte) *Continent {
    model.data = data
    return &model
}

func (model *Continent) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}