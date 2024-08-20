package models

import (
    "encoding/json"
    "errors"
)

// TemplateVariable Model
type TemplateVariable struct {
    // Variable Name.
    Name string `json:"name"`
    // Variable Description.
    Description string `json:"description"`
    // Variable Value.
    Value string `json:"value"`
    // Variable Placeholder.
    Placeholder string `json:"placeholder"`
    // Is the variable required?
    Required bool `json:"required"`
    // Variable Type.
    Type string `json:"xtype"`

    // Used by Decode() method
    data []byte
}

func (model TemplateVariable) New(data []byte) *TemplateVariable {
    model.data = data
    return &model
}

func (model *TemplateVariable) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}