package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasskey Model
type PolicyPasskey struct {
	// Policy ID.
	Id string `json:"$id"`
	// Relying party ID passkeys are bound to. Empty until configured.
	RpId string `json:"rpId"`
	// Web origins allowed to register and sign in with passkeys.
	Origins []string `json:"origins"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasskey) New(data []byte) *PolicyPasskey {
	model.data = data
	return &model
}

func (model *PolicyPasskey) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
