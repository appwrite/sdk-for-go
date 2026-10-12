package models

import (
	"encoding/json"
	"errors"
)

// Passkey Model
type Passkey struct {
	// Passkey ID.
	Id string `json:"$id"`
	// Passkey creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Passkey update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Passkey name.
	Name string `json:"name"`
	// Most recent sign-in with this passkey in ISO 8601 format. Empty until the
	// passkey is used to sign in.
	AccessedAt string `json:"accessedAt"`

	// Used by Decode() method
	data []byte
}

func (model Passkey) New(data []byte) *Passkey {
	model.data = data
	return &model
}

func (model *Passkey) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
