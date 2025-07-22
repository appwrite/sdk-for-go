package models

import (
    "encoding/json"
    "errors"
)

// ColumnBoolean Model
type ColumnBoolean struct {
    // Column Key.
    Key string `json:"key"`
    // Column type.
    Type string `json:"xtype"`
    // Column status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Error message. Displays error generated on failure of creating or deleting
    // an column.
    Error string `json:"xerror"`
    // Is column required?
    Required bool `json:"required"`
    // Is column an array?
    Array bool `json:"array"`
    // Column creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Column update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Default value for attribute when not provided. Cannot be set when attribute
    // is required.
    Default bool `json:"xdefault"`

    // Used by Decode() method
    data []byte
}

func (model ColumnBoolean) New(data []byte) *ColumnBoolean {
    model.data = data
    return &model
}

func (model *ColumnBoolean) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}