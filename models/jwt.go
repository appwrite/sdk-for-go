package models

import (
    "encoding/json"
    "errors"
)

// JWT Model
type Jwt struct {
    // JWT encoded string.
    Jwt string `json:"jwt"`

    // Used by Decode() method
    data []byte
}

func (model Jwt) New(data []byte) *Jwt {
    model.data = data
    return &model
}

func (model *Jwt) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}