package models

import (
    "encoding/json"
    "errors"
)

// Document Model
type Document struct {
    // Document ID.
    Id string `json:"$id"`
    // Collection ID.
    CollectionId string `json:"$collectionId"`
    // Database ID.
    DatabaseId string `json:"$databaseId"`
    // Document creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Document update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Document permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`

    // Used by Decode() method
    data []byte
}

func (model Document) New(data []byte) *Document {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *Document) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
