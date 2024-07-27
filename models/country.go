package models

import (
    "encoding/json"
    "errors"
)

// Country Model
type Country struct {
    // Country name.
    Name string `json:"name"`
    // Country two-character ISO 3166-1 alpha code.
    Code string `json:"code"`

    // Used by Decode() method
    data []byte
}

func (model Country) New(data []byte) *Country {
    model.data = data
    return &model
}

func (model *Country) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}