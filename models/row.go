package models

import (
    "encoding/json"
    "errors"
)

// Row Model
type Row struct {
    // Row ID.
    Id string `json:"$id"`
    // Row automatically incrementing ID.
    Sequence int `json:"$sequence"`
    // Table ID.
    TableId string `json:"$tableId"`
    // Database ID.
    DatabaseId string `json:"$databaseId"`
    // Row creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Row update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Row permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`

    // Used by Decode() method
    data []byte
}

func (model Row) New(data []byte) *Row {
    model.data = data
    return &model
}

// Use this method to get response in desired type
func (model *Row) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}