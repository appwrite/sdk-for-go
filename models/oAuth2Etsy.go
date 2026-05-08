package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Etsy Model
type OAuth2Etsy struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Etsy OAuth2 keystring.
    KeyString string `json:"keyString"`
    // Etsy OAuth2 shared secret.
    SharedSecret string `json:"sharedSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Etsy) New(data []byte) *OAuth2Etsy {
    model.data = data
    return &model
}

func (model *OAuth2Etsy) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}