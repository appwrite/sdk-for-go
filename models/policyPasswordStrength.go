package models

import (
	"encoding/json"
	"errors"
)

// PolicyPasswordStrength Model
type PolicyPasswordStrength struct {
	// Policy ID.
	Id string `json:"$id"`
	// Minimum password length required for user passwords.
	Min int `json:"min"`
	// Whether passwords must include at least one uppercase letter.
	Uppercase bool `json:"uppercase"`
	// Whether passwords must include at least one lowercase letter.
	Lowercase bool `json:"lowercase"`
	// Whether passwords must include at least one number.
	Number bool `json:"number"`
	// Whether passwords must include at least one symbol.
	Symbols bool `json:"symbols"`

	// Used by Decode() method
	data []byte
}

func (model PolicyPasswordStrength) New(data []byte) *PolicyPasswordStrength {
	model.data = data
	return &model
}

func (model *PolicyPasswordStrength) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
