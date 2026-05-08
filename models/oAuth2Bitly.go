package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Bitly Model
type OAuth2Bitly struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Bitly OAuth2 client ID.
    ClientId string `json:"clientId"`
    // Bitly OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Bitly) New(data []byte) *OAuth2Bitly {
    model.data = data
    return &model
}

func (model *OAuth2Bitly) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}