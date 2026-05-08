package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Dropbox Model
type OAuth2Dropbox struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Dropbox OAuth2 app key.
    AppKey string `json:"appKey"`
    // Dropbox OAuth2 app secret.
    AppSecret string `json:"appSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Dropbox) New(data []byte) *OAuth2Dropbox {
    model.data = data
    return &model
}

func (model *OAuth2Dropbox) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}