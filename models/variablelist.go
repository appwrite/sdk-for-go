package models

import (
    "encoding/json"
    "errors"
)

// VariablesList Model
type VariableList struct {
    // Total number of variables documents that matched your query.
    Total int `json:"total"`
    // List of variables.
    Variables []Variable `json:"variables"`

    // Used by Decode() method
    data []byte
}

func (model VariableList) New(data []byte) *VariableList {
    model.data = data
    return &model
}

func (model *VariableList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}