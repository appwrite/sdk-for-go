package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Stripe Model
type OAuth2Stripe struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Stripe OAuth2 client ID.
    ClientId string `json:"clientId"`
    // Stripe OAuth2 API secret key.
    ApiSecretKey string `json:"apiSecretKey"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Stripe) New(data []byte) *OAuth2Stripe {
    model.data = data
    return &model
}

func (model *OAuth2Stripe) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}