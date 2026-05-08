package models

import (
    "encoding/json"
    "errors"
)

// OAuth2X Model
type OAuth2X struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // X OAuth2 customer key.
    CustomerKey string `json:"customerKey"`
    // X OAuth2 secret key.
    SecretKey string `json:"secretKey"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2X) New(data []byte) *OAuth2X {
    model.data = data
    return &model
}

func (model *OAuth2X) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}