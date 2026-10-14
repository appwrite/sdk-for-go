package models

import (
	"encoding/json"
	"errors"
)

// MFAFactors Model
type MfaFactors struct {
	// Can TOTP be used for MFA challenge for this account.
	Totp bool `json:"totp"`
	// Can phone (SMS) be used for MFA challenge for this account.
	Phone bool `json:"phone"`
	// Can email be used for MFA challenge for this account.
	Email bool `json:"email"`
	// Can recovery code be used for MFA challenge for this account.
	RecoveryCode bool `json:"recoveryCode"`
	// Can custom factor be used for MFA challenge for this account.
	Custom bool `json:"custom"`

	// Used by Decode() method
	data []byte
}

func (model MfaFactors) New(data []byte) *MfaFactors {
	model.data = data
	return &model
}

func (model *MfaFactors) UnmarshalJSON(data []byte) error {
	type alias MfaFactors
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"totp\":false,\"phone\":false,\"email\":false,\"recoveryCode\":false,\"custom\":false}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
}

func (model *MfaFactors) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
