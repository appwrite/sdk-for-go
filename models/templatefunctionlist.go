package models

import (
    "encoding/json"
    "errors"
)

// FunctionTemplatesList Model
type TemplateFunctionList struct {
    // Total number of templates documents that matched your query.
    Total int `json:"total"`
    // List of templates.
    Templates []TemplateFunction `json:"templates"`

    // Used by Decode() method
    data []byte
}

func (model TemplateFunctionList) New(data []byte) *TemplateFunctionList {
    model.data = data
    return &model
}

func (model *TemplateFunctionList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}