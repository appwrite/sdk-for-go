package models

import (
    "encoding/json"
    "errors"
)

// FrameworksList Model
type FrameworkList struct {
    // Total number of frameworks documents that matched your query.
    Total int `json:"total"`
    // List of frameworks.
    Frameworks []Framework `json:"frameworks"`

    // Used by Decode() method
    data []byte
}

func (model FrameworkList) New(data []byte) *FrameworkList {
    model.data = data
    return &model
}

func (model *FrameworkList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}