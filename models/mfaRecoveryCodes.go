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

func (model *MfaRecoveryCodes) UnmarshalJSON(data []byte) error {
	type alias MfaRecoveryCodes
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"recoveryCodes\":[]}"), &fields); err != nil {
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
