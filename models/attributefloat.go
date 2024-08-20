package models

import (
    "encoding/json"
    "errors"
)

// AttributeFloat Model
type AttributeFloat struct {
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
    // Minimum value to enforce for new documents.
    Min float64 `json:"min"`
    // Maximum value to enforce for new documents.
    Max float64 `json:"max"`
    // Default value for attribute when not provided. Cannot be set when attribute
    // is required.
    Default float64 `json:"xdefault"`

    // Used by Decode() method
    data []byte
}

func (model AttributeFloat) New(data []byte) *AttributeFloat {
    model.data = data
    return &model
}

func (model *AttributeFloat) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}