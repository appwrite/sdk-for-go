package models

import (
    "encoding/json"
    "errors"
)

// MFARecoveryCodes Model
type MfaRecoveryCodes struct {
    // Recovery codes.
    RecoveryCodes []string `json:"recoveryCodes"`

    // Used by Decode() method
    data []byte
}

func (model MfaRecoveryCodes) New(data []byte) *MfaRecoveryCodes {
    model.data = data
    return &model
}

func (model *MfaRecoveryCodes) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
