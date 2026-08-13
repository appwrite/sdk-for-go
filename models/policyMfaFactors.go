package models

import (
    "encoding/json"
    "errors"
)

// PolicyMFAFactors Model
type PolicyMfaFactors struct {
    // Policy ID.
    Id string `json:"$id"`
    // Whether TOTP can be used to complete an MFA challenge.
    Totp bool `json:"totp"`
    // Whether email can be used to complete an MFA challenge.
    Email bool `json:"email"`
    // Whether phone (SMS) can be used to complete an MFA challenge.
    Phone bool `json:"phone"`
    // Whether the custom factor can be used to complete an MFA challenge.
    Custom bool `json:"custom"`

    // Used by Decode() method
    data []byte
}

func (model PolicyMfaFactors) New(data []byte) *PolicyMfaFactors {
    model.data = data
    return &model
}

func (model *PolicyMfaFactors) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}