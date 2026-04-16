package models

import (
    "encoding/json"
    "errors"
)

// Key Model
type Key struct {
    // Key ID.
    Id string `json:"$id"`
    // Key creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Key update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Key name.
    Name string `json:"name"`
    // Key expiration date in ISO 8601 format.
    Expire string `json:"expire"`
    // Allowed permission scopes.
    Scopes []string `json:"scopes"`
    // Secret key.
    Secret string `json:"secret"`
    // Most recent access date in ISO 8601 format. This attribute is only updated
    // again after 24 hours.
    AccessedAt string `json:"accessedAt"`
    // List of SDK user agents that used this key.
    Sdks []string `json:"sdks"`

    // Used by Decode() method
    data []byte
}

func (model Key) New(data []byte) *Key {
    model.data = data
    return &model
}

func (model *Key) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}