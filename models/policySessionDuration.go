package models

import (
    "encoding/json"
    "errors"
)

// PolicySessionDuration Model
type PolicySessionDuration struct {
    // Policy ID.
    Id string `json:"$id"`
    // Session duration in seconds.
    Duration int `json:"duration"`

    // Used by Decode() method
    data []byte
}

func (model PolicySessionDuration) New(data []byte) *PolicySessionDuration {
    model.data = data
    return &model
}

func (model *PolicySessionDuration) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}