package models

import (
    "encoding/json"
    "errors"
)

// Provider Model
type Provider struct {
    // Provider ID.
    Id string `json:"$id"`
    // Provider creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Provider update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // The name for the provider instance.
    Name string `json:"name"`
    // The name of the provider service.
    Provider string `json:"provider"`
    // Is provider enabled?
    Enabled bool `json:"enabled"`
    // Type of provider.
    Type string `json:"xtype"`
    // Provider credentials.
    Credentials interface{} `json:"credentials"`
    // Provider options.
    Options interface{} `json:"options"`

    // Used by Decode() method
    data []byte
}

func (model Provider) New(data []byte) *Provider {
    model.data = data
    return &model
}

func (model *Provider) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}