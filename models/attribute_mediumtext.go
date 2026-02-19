package models

import (
    "encoding/json"
    "errors"
)

// AttributeMediumtext Model
type AttributeMediumtext struct {
    // Attribute Key.
    Key string `json:"key"`
    // Attribute type.
    Type string `json:"type"`
    // Attribute status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Error message. Displays error generated on failure of creating or deleting
    // an attribute.
    Error string `json:"error"`
    // Is attribute required?
    Required bool `json:"required"`
    // Is attribute an array?
    Array bool `json:"array"`
    // Attribute creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Attribute update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Default value for attribute when not provided. Cannot be set when attribute
    // is required.
    Default string `json:"default"`
    // Defines whether this attribute is encrypted or not.
    Encrypt bool `json:"encrypt"`

    // Used by Decode() method
    data []byte
}

func (model AttributeMediumtext) New(data []byte) *AttributeMediumtext {
    model.data = data
    return &model
}

func (model *AttributeMediumtext) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}