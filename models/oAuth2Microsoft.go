package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Microsoft Model
type OAuth2Microsoft struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Microsoft OAuth2 application ID.
    ApplicationId string `json:"applicationId"`
    // Microsoft OAuth2 application secret.
    ApplicationSecret string `json:"applicationSecret"`
    // Microsoft Entra ID tenant identifier. Use 'common', 'organizations',
    // 'consumers' or a specific tenant ID.
    Tenant string `json:"tenant"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Microsoft) New(data []byte) *OAuth2Microsoft {
    model.data = data
    return &model
}

func (model *OAuth2Microsoft) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}