package models

import (
    "encoding/json"
    "errors"
)

// MFAFactors Model
type MfaFactors struct {
    // Can TOTP be used for MFA challenge for this account.
    Totp bool `json:"totp"`
    // Can phone (SMS) be used for MFA challenge for this account.
    Phone bool `json:"phone"`
    // Can email be used for MFA challenge for this account.
    Email bool `json:"email"`
    // Can recovery code be used for MFA challenge for this account.
    RecoveryCode bool `json:"recoveryCode"`

    // Used by Decode() method
    data []byte
}

func (model MfaFactors) New(data []byte) *MfaFactors {
    model.data = data
    return &model
}

func (model *MfaFactors) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
