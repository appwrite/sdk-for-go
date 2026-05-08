package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Disqus Model
type OAuth2Disqus struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Disqus OAuth2 public key.
    PublicKey string `json:"publicKey"`
    // Disqus OAuth2 secret key.
    SecretKey string `json:"secretKey"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Disqus) New(data []byte) *OAuth2Disqus {
    model.data = data
    return &model
}

func (model *OAuth2Disqus) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}