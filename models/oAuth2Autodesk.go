package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Autodesk Model
type OAuth2Autodesk struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Autodesk OAuth2 client ID.
    ClientId string `json:"clientId"`
    // Autodesk OAuth2 client secret.
    ClientSecret string `json:"clientSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Autodesk) New(data []byte) *OAuth2Autodesk {
    model.data = data
    return &model
}

func (model *OAuth2Autodesk) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}