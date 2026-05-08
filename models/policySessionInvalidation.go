package models

import (
    "encoding/json"
    "errors"
)

// PolicySessionInvalidation Model
type PolicySessionInvalidation struct {
    // Policy ID.
    Id string `json:"$id"`
    // Whether session invalidation policy is enabled.
    Enabled bool `json:"enabled"`

    // Used by Decode() method
    data []byte
}

func (model PolicySessionInvalidation) New(data []byte) *PolicySessionInvalidation {
    model.data = data
    return &model
}

func (model *PolicySessionInvalidation) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}