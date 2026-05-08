package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Zoho Model
type OAuth2Zoho struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Zoho OAuth2 client ID.
    ClientId string `json:"clientId"`
    // Zoho OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Zoho) New(data []byte) *OAuth2Zoho {
    model.data = data
    return &model
}

func (model *OAuth2Zoho) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}