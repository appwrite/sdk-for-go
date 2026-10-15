package models

import (
	"encoding/json"
	"errors"
)

// MFAType Model
type MfaType struct {
	// Secret token used for TOTP factor.
	Secret string `json:"secret"`
	// URI for authenticator apps.
	Uri string `json:"uri"`

	// Used by Decode() method
	data []byte
}

func (model MfaType) New(data []byte) *MfaType {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *MfaType) UnmarshalJSON(data []byte) error {
	type alias MfaType
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Secret json.RawMessage `json:"secret"`
		Uri    json.RawMessage `json:"uri"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}
	if present.Uri == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Uri); err != nil {
			return err
		}
	}

	*model = MfaType(decoded)
	return nil
}

func (model *MfaType) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
