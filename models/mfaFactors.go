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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *MfaFactors) UnmarshalJSON(data []byte) error {
	type alias MfaFactors
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Totp         json.RawMessage `json:"totp"`
		Phone        json.RawMessage `json:"phone"`
		Email        json.RawMessage `json:"email"`
		RecoveryCode json.RawMessage `json:"recoveryCode"`
		Custom       json.RawMessage `json:"custom"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Totp == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Totp); err != nil {
			return err
		}
	}
	if present.Phone == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Phone); err != nil {
			return err
		}
	}
	if present.Email == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Email); err != nil {
			return err
		}
	}
	if present.RecoveryCode == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.RecoveryCode); err != nil {
			return err
		}
	}
	if present.Custom == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Custom); err != nil {
			return err
		}
	}

	*model = MfaFactors(decoded)
	return nil
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
