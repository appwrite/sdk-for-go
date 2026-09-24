package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasswordPwned Model
type PolicyPasswordPwned struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether passwords are checked against known data breaches and the result
	// recorded on the user.
	Enabled bool `json:"enabled"`
	// Whether a sign-in with a breached password is refused until the password is
	// reset.
	Sessions bool `json:"sessions"`
	// Whether a breached password is rejected when a user signs up or sets a new
	// password.
	Users bool `json:"users"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasswordPwned) New(data []byte) *PolicyPasswordPwned {
	model.data = data
	return &model
}

func (model *PolicyPasswordPwned) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
