package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Notion Model
type OAuth2Notion struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Notion OAuth2 client ID.
    OauthClientId string `json:"oauthClientId"`
    // Notion OAuth2 client secret.
    OauthClientSecret string `json:"oauthClientSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Notion) New(data []byte) *OAuth2Notion {
    model.data = data
    return &model
}

func (model *OAuth2Notion) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}