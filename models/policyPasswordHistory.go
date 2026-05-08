package models

import (
    "encoding/json"
    "errors"
)

// PolicyPasswordHistory Model
type PolicyPasswordHistory struct {
    // Policy ID.
    Id string `json:"$id"`
    // Password history length. A value of 0 means the policy is disabled.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model PolicyPasswordHistory) New(data []byte) *PolicyPasswordHistory {
    model.data = data
    return &model
}

func (model *PolicyPasswordHistory) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}