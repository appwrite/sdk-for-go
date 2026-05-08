package models

import (
    "encoding/json"
    "errors"
)

// PolicySessionLimit Model
type PolicySessionLimit struct {
    // Policy ID.
    Id string `json:"$id"`
    // Maximum number of sessions allowed per user. A value of 0 means the policy
    // is disabled.
    Total int `json:"total"`

    // Used by Decode() method
    data []byte
}

func (model PolicySessionLimit) New(data []byte) *PolicySessionLimit {
    model.data = data
    return &model
}

func (model *PolicySessionLimit) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}