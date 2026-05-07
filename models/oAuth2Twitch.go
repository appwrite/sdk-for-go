package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Twitch Model
type OAuth2Twitch struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Twitch OAuth2 client ID.
    ClientId string `json:"clientId"`
    // Twitch OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Twitch) New(data []byte) *OAuth2Twitch {
    model.data = data
    return &model
}

func (model *OAuth2Twitch) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}