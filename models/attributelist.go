package models

import (
    "encoding/json"
    "errors"
)

// AttributesList Model
type AttributeList struct {
    // Total number of attributes in the given collection.
    Total int `json:"total"`
    // List of attributes.
    Attributes []map[string]any `json:"attributes"`

    // Used by Decode() method
    data []byte
}

func (model AttributeList) New(data []byte) *AttributeList {
    model.data = data
    return &model
}

func (model *AttributeList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
