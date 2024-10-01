package models

import (
    "encoding/json"
    "errors"
)

// MFAChallenge Model
type MfaChallenge struct {
    // Token ID.
    Id string `json:"$id"`
    // Token creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // User ID.
    UserId string `json:"userId"`
    // Token expiration date in ISO 8601 format.
    Expire string `json:"expire"`

    // Used by Decode() method
    data []byte
}

func (model MfaChallenge) New(data []byte) *MfaChallenge {
    model.data = data
    return &model
}

func (model *MfaChallenge) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
