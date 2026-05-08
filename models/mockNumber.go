package models

import (
    "encoding/json"
    "errors"
)

// MockNumber Model
type MockNumber struct {
    // Mock phone number for testing phone authentication. Useful for testing
    // phone authentication without sending an SMS.
    Number string `json:"number"`
    // Mock OTP for the number.
    Otp string `json:"otp"`
    // Attribute creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Attribute update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`

    // Used by Decode() method
    data []byte
}

func (model MockNumber) New(data []byte) *MockNumber {
    model.data = data
    return &model
}

func (model *MockNumber) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}