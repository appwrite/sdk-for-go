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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyPasswordPwned) UnmarshalJSON(data []byte) error {
	type alias PolicyPasswordPwned
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id       json.RawMessage `json:"$id"`
		Enabled  json.RawMessage `json:"enabled"`
		Sessions json.RawMessage `json:"sessions"`
		Users    json.RawMessage `json:"users"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Sessions == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Sessions); err != nil {
			return err
		}
	}
	if present.Users == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Users); err != nil {
			return err
		}
	}

	*model = PolicyPasswordPwned(decoded)
	return nil
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
