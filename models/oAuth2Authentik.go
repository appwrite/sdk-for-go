package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Authentik Model
type OAuth2Authentik struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Authentik OAuth2 client ID.
    ClientId string `json:"clientId"`
    // Authentik OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`
    // Authentik OAuth2 endpoint domain.
    Endpoint string `json:"endpoint"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Authentik) New(data []byte) *OAuth2Authentik {
    model.data = data
    return &model
}

func (model *OAuth2Authentik) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}