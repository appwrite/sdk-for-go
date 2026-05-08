package models

import (
    "encoding/json"
    "errors"
)

// OAuth2FusionAuth Model
type OAuth2FusionAuth struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // FusionAuth OAuth2 client ID.
    ClientId string `json:"clientId"`
    // FusionAuth OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`
    // FusionAuth OAuth2 endpoint domain.
    Endpoint string `json:"endpoint"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2FusionAuth) New(data []byte) *OAuth2FusionAuth {
    model.data = data
    return &model
}

func (model *OAuth2FusionAuth) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}