package models

import (
    "encoding/json"
    "errors"
)

// Table Model
type Table struct {
    // Table ID.
    Id string `json:"$id"`
    // Table creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Table update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Table permissions. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    Permissions []string `json:"$permissions"`
    // Database ID.
    DatabaseId string `json:"databaseId"`
    // Table name.
    Name string `json:"name"`
    // Table enabled. Can be 'enabled' or 'disabled'. When disabled, the table is
    // inaccessible to users, but remains accessible to Server SDKs using API
    // keys.
    Enabled bool `json:"enabled"`
    // Whether row-level permissions are enabled. [Learn more about
    // permissions](https://appwrite.io/docs/permissions).
    RowSecurity bool `json:"rowSecurity"`
    // Table columns.
    Columns []interface{} `json:"columns"`
    // Table indexes.
    Indexes []ColumnIndex `json:"indexes"`

    // Used by Decode() method
    data []byte
}

func (model Table) New(data []byte) *Table {
    model.data = data
    return &model
}

func (model *Table) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}