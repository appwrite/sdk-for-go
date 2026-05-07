package models

import (
    "encoding/json"
    "errors"
)

// OAuth2Salesforce Model
type OAuth2Salesforce struct {
    // OAuth2 provider ID.
    Id string `json:"$id"`
    // OAuth2 provider is active and can be used to create sessions.
    Enabled bool `json:"enabled"`
    // Salesforce OAuth2 consumer key.
    CustomerKey string `json:"customerKey"`
    // Salesforce OAuth2 consumer secret.
    CustomerSecret string `json:"customerSecret"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2Salesforce) New(data []byte) *OAuth2Salesforce {
    model.data = data
    return &model
}

func (model *OAuth2Salesforce) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}