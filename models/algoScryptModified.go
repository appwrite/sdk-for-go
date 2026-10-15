package models

import (
	"encoding/json"
	"errors"
)

// AlgoScryptModified Model
type AlgoScryptModified struct {
	// Algo type.
	Type string `json:"type"`
	// Salt used to compute hash.
	Salt string `json:"salt"`
	// Separator used to compute hash.
	SaltSeparator string `json:"saltSeparator"`
	// Key used to compute hash.
	SignerKey string `json:"signerKey"`

	// Used by Decode() method
	data []byte
}

func (model AlgoScryptModified) New(data []byte) *AlgoScryptModified {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AlgoScryptModified) UnmarshalJSON(data []byte) error {
	type alias AlgoScryptModified
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Type          json.RawMessage `json:"type"`
		Salt          json.RawMessage `json:"salt"`
		SaltSeparator json.RawMessage `json:"saltSeparator"`
		SignerKey     json.RawMessage `json:"signerKey"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"scryptMod\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Salt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Salt); err != nil {
			return err
		}
	}
	if present.SaltSeparator == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SaltSeparator); err != nil {
			return err
		}
	}
	if present.SignerKey == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SignerKey); err != nil {
			return err
		}
	}

	*model = AlgoScryptModified(decoded)
	return nil
}

func (model *AlgoScryptModified) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
