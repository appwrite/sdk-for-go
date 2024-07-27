package models

import (
    "encoding/json"
    "errors"
)

// Collection Model
type Collection struct {
    // Collection ID.
    Id string `json:"$id"`
    // Collection creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Collection update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Collection permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`
    // Database ID.
    DatabaseId string `json:"databaseId"`
    // Collection name.
    Name string `json:"name"`
    // Collection enabled. Can be 'enabled' or 'disabled'. When disabled, the
    // collection is inaccessible to users, but remains accessible to Server SDKs
    // using API keys.
    Enabled bool `json:"enabled"`
    // Whether document-level permissions are enabled. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    DocumentSecurity bool `json:"documentSecurity"`
    // Collection attributes.
    Attributes []string `json:"attributes"`
    // Collection indexes.
    Indexes []Index `json:"indexes"`

    // Used by Decode() method
    data []byte
}

func (model Collection) New(data []byte) *Collection {
    model.data = data
    return &model
}

func (model *Collection) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}