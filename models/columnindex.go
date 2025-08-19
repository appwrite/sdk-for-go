package models

import (
    "encoding/json"
    "errors"
)

// Index Model
type ColumnIndex struct {
    // Index ID.
    Id string `json:"$id"`
    // Index creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Index update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Index Key.
    Key string `json:"key"`
    // Index type.
    Type string `json:"xtype"`
    // Index status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Error message. Displays error generated on failure of creating or deleting
    // an index.
    Error string `json:"xerror"`
    // Index columns.
    Columns []string `json:"columns"`
    // Index columns length.
    Lengths []int `json:"lengths"`
    // Index orders.
    Orders []string `json:"orders"`

    // Used by Decode() method
    data []byte
}

func (model ColumnIndex) New(data []byte) *ColumnIndex {
    model.data = data
    return &model
}

func (model *ColumnIndex) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}