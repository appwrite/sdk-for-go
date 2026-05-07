package models

import (
    "encoding/json"
    "errors"
)

// PolicySessionAlert Model
type PolicySessionAlert struct {
    // Policy ID.
    Id string `json:"$id"`
    // Whether session alert policy is enabled.
    Enabled bool `json:"enabled"`

    // Used by Decode() method
    data []byte
}

func (model PolicySessionAlert) New(data []byte) *PolicySessionAlert {
    model.data = data
    return &model
}

func (model *PolicySessionAlert) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}