package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Dailymotion Model
type OAuth2Dailymotion struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Dailymotion OAuth2 API key.
    ApiKey string `json:"apiKey"`
    // Dailymotion OAuth2 API secret.
    ApiSecret string `json:"apiSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Dailymotion) New(data []byte) *OAuth2Dailymotion {
    model.data = data
    return &model
}

func (model *OAuth2Dailymotion) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}