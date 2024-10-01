package models

import (
    "encoding/json"
    "errors"
)

// Target Model
type Target struct {
    // Target ID.
    Id string `json:"$id"`
    // Target creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Target update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Target Name.
    Name string `json:"name"`
    // User ID.
    UserId string `json:"userId"`
    // Provider ID.
    ProviderId string `json:"providerId"`
    // The target provider type. Can be one of the following: `email`, `sms` or
    // `push`.
    ProviderType string `json:"providerType"`
    // The target identifier.
    Identifier string `json:"identifier"`

    // Used by Decode() method
    data []byte
}

func (model Target) New(data []byte) *Target {
    model.data = data
    return &model
}

func (model *Target) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
