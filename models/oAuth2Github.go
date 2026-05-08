package models

import (
    "encoding/json"
    "errors"
)

// OAuth2GitHub Model
type OAuth2Github struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // GitHub OAuth2 client ID. For GitHub Apps, use the "App ID" when both an App
    // ID and client ID are available.
    ClientId string `json:"clientId"`
    // GitHub OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Github) New(data []byte) *OAuth2Github {
    model.data = data
    return &model
}

func (model *OAuth2Github) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}