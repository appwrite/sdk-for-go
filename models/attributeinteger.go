package models

import (
    "encoding/json"
    "errors"
)

// AttributeInteger Model
type AttributeInteger struct {
    // Attribute Key.
    Key string `json:"key"`
    // Attribute type.
    Type string `json:"xtype"`
    // Attribute status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Error message. Displays error generated on failure of creating or deleting
    // an attribute.
    Error string `json:"xerror"`
    // Is attribute required?
    Required bool `json:"required"`
    // Is attribute an array?
    Array bool `json:"array"`
    // Attribute creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Attribute update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Minimum value to enforce for new documents.
    Min int `json:"min"`
    // Maximum value to enforce for new documents.
    Max int `json:"max"`
    // Default value for attribute when not provided. Cannot be set when attribute
    // is required.
    Default int `json:"xdefault"`

    // Used by Decode() method
    data []byte
}

func (model AttributeInteger) New(data []byte) *AttributeInteger {
    model.data = data
    return &model
}

func (model *AttributeInteger) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}