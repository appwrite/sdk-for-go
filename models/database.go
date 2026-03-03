package models

import (
    "encoding/json"
    "errors"
)

// Database Model
type Database struct {
    // Database ID.
    Id string `json:"$id"`
    // Database name.
    Name string `json:"name"`
    // Database creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Database update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // If database is enabled. Can be 'enabled' or 'disabled'. When disabled, the
    // database is inaccessible to users, but remains accessible to Server SDKs
    // using API keys.
    Enabled bool `json:"enabled"`
    // Database type.
    Type string `json:"type"`

    // Used by Decode() method
    data []byte
}

func (model Database) New(data []byte) *Database {
    model.data = data
    return &model
}

func (model *Database) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}