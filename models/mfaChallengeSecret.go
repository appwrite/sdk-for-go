package models

import (
    "encoding/json"
    "errors"
)

// MFAChallengeSecret Model
type MfaChallengeSecret struct {
    // Token ID.
    Id string `json:"$id"`
    // Token creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // User ID.
    UserId string `json:"userId"`
    // Token expiration date in ISO 8601 format.
    Expire string `json:"expire"`
    // Challenge code to be delivered to the end user through a custom channel.
    Code string `json:"code"`

    // Used by Decode() method
    data []byte
}

func (model MfaChallengeSecret) New(data []byte) *MfaChallengeSecret {
    model.data = data
    return &model
}

func (model *MfaChallengeSecret) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}