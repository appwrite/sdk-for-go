package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Tradeshift Model
type OAuth2Tradeshift struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Tradeshift OAuth2 client ID.
    Oauth2ClientId string `json:"oauth2ClientId"`
    // Tradeshift OAuth2 client secret.
    Oauth2ClientSecret string `json:"oauth2ClientSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Tradeshift) New(data []byte) *OAuth2Tradeshift {
    model.data = data
    return &model
}

func (model *OAuth2Tradeshift) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}