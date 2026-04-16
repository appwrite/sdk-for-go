package models

import (
    "encoding/json"
    "errors"
)

// ColumnMediumtext Model
type ColumnMediumtext struct {
    // Column Key.
    Key string `json:"key"`
    // Column type.
    Type string `json:"type"`
    // Column status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Error message. Displays error generated on failure of creating or deleting
    // an column.
    Error string `json:"error"`
    // Is column required?
    Required bool `json:"required"`
    // Is column an array?
    Array bool `json:"array"`
    // Column creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Column update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Default value for column when not provided. Cannot be set when column is
    // required.
    Default string `json:"default"`
    // Defines whether this column is encrypted or not.
    Encrypt bool `json:"encrypt"`

    // Used by Decode() method
    data []byte
}

func (model ColumnMediumtext) New(data []byte) *ColumnMediumtext {
    model.data = data
    return &model
}

func (model *ColumnMediumtext) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}