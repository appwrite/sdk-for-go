package models

import (
    "encoding/json"
    "errors"
)

// PolicyDenyCorporateEmail Model
type PolicyDenyCorporateEmail struct {
    // Policy ID.
    Id string `json:"$id"`
    // Whether the deny non-corporate email policy is enabled.
    Enabled bool `json:"enabled"`

    // Used by Decode() method
    data []byte
}

func (model PolicyDenyCorporateEmail) New(data []byte) *PolicyDenyCorporateEmail {
    model.data = data
    return &model
}

func (model *PolicyDenyCorporateEmail) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}