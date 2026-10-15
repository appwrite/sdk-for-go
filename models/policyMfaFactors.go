package models

import (
	"encoding/json"
	"errors"
)

// PolicyMFAFactors Model
type PolicyMfaFactors struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether TOTP can be used to complete an MFA challenge.
	Totp bool `json:"totp"`
	// Whether email can be used to complete an MFA challenge.
	Email bool `json:"email"`
	// Whether phone (SMS) can be used to complete an MFA challenge.
	Phone bool `json:"phone"`
	// Whether the custom factor can be used to complete an MFA challenge.
	Custom bool `json:"custom"`

	// Used by Decode() method
	data []byte
}

func (model PolicyMfaFactors) New(data []byte) *PolicyMfaFactors {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PolicyMfaFactors) UnmarshalJSON(data []byte) error {
	type alias PolicyMfaFactors
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id     json.RawMessage `json:"$id"`
		Totp   json.RawMessage `json:"totp"`
		Email  json.RawMessage `json:"email"`
		Phone  json.RawMessage `json:"phone"`
		Custom json.RawMessage `json:"custom"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Totp == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Totp); err != nil {
			return err
		}
	}
	if present.Email == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Email); err != nil {
			return err
		}
	}
	if present.Phone == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Phone); err != nil {
			return err
		}
	}
	if present.Custom == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Custom); err != nil {
			return err
		}
	}

	*model = PolicyMfaFactors(decoded)
	return nil
}

func (model *PolicyMfaFactors) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
