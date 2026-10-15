package models

import (
	"encoding/json"
	"errors"
)

// MFARecoveryCodes Model
type MfaRecoveryCodes struct {
	// Recovery codes.
	RecoveryCodes []string `json:"recoveryCodes"`

	// Used by Decode() method
	data []byte
}

func (model MfaRecoveryCodes) New(data []byte) *MfaRecoveryCodes {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *MfaRecoveryCodes) UnmarshalJSON(data []byte) error {
	type alias MfaRecoveryCodes
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		RecoveryCodes json.RawMessage `json:"recoveryCodes"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.RecoveryCodes == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.RecoveryCodes); err != nil {
			return err
		}
	}

	*model = MfaRecoveryCodes(decoded)
	return nil
}

func (model *MfaRecoveryCodes) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
