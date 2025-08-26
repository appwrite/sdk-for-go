package models

import (
    "encoding/json"
    "errors"
)

// ColumnEnum Model
type ColumnEnum struct {
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
    // Array of elements in enumerated type.
    Elements []string `json:"elements"`
    // String format.
    Format string `json:"format"`
    // Default value for column when not provided. Cannot be set when column is
    // required.
    Default string `json:"default"`

    // Used by Decode() method
    data []byte
}

func (model ColumnEnum) New(data []byte) *ColumnEnum {
    model.data = data
    return &model
}

func (model *ColumnEnum) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}