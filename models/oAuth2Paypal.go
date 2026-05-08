package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Paypal Model
type OAuth2Paypal struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // PayPal OAuth2 client ID.
    ClientId string `json:"clientId"`
    // PayPal OAuth2 secret key.
    SecretKey string `json:"secretKey"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Paypal) New(data []byte) *OAuth2Paypal {
    model.data = data
    return &model
}

func (model *OAuth2Paypal) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}