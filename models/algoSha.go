package models

import (
	"encoding/json"
	"errors"
)

// AlgoSHA Model
type AlgoSha struct {
	// Algo type.
	Type string `json:"type"`

	// Used by Decode() method
	data []byte
}

func (model AlgoSha) New(data []byte) *AlgoSha {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AlgoSha) UnmarshalJSON(data []byte) error {
	type alias AlgoSha
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Type json.RawMessage `json:"type"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"sha\""), &decoded.Type); err != nil {
			return err
		}
	}

	*model = AlgoSha(decoded)
	return nil
}

func (model *AlgoSha) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
