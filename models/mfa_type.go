package models

import (
    "encoding/json"
    "errors"
)

// MFAType Model
type MfaType struct {
    // Secret token used for TOTP factor.
    Secret string `json:"secret"`
    // URI for authenticator apps.
    Uri string `json:"uri"`

    // Used by Decode() method
    data []byte
}

func (model MfaType) New(data []byte) *MfaType {
    model.data = data
    return &model
}

func (model *MfaType) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}