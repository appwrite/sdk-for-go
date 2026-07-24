package models

import (
    "encoding/json"
    "errors"
)

// AppScopesList Model
type AppScopeList struct {
    // Total number of scopes that matched your query.
    Total int `json:"total"`
    // List of scopes.
    Scopes []AppScope `json:"scopes"`

    // Used by Decode() method
    data []byte
}

func (model AppScopeList) New(data []byte) *AppScopeList {
    model.data = data
    return &model
}

func (model *AppScopeList) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}