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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *MockNumber) UnmarshalJSON(data []byte) error {
	type alias MockNumber
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Number    json.RawMessage `json:"number"`
		Otp       json.RawMessage `json:"otp"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UpdatedAt json.RawMessage `json:"$updatedAt"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Number == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Number); err != nil {
			return err
		}
	}
	if present.Otp == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Otp); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}

	*model = MockNumber(decoded)
	return nil
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
