package models

import (
    "encoding/json"
    "errors"
)

// OAuth2ProvidersList Model
type OAuth2ProviderList struct {
    // Total number of OAuth2 providers in the given project.
    Total int `json:"total"`
    // List of OAuth2 providers.
    Providers []interface{} `json:"providers"`

    // Used by Decode() method
    data []byte
}

func (model OAuth2ProviderList) New(data []byte) *OAuth2ProviderList {
    model.data = data
    return &model
}

func (model *OAuth2ProviderList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}