package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Amazon Model
type OAuth2Amazon struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Amazon OAuth2 client ID.
    ClientId string `json:"clientId"`
    // Amazon OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Amazon) New(data []byte) *OAuth2Amazon {
    model.data = data
    return &model
}

func (model *OAuth2Amazon) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}