package models

import (
	"encoding/json"
	"errors"
)

// AlgoPHPass Model
type AlgoPhpass struct {
	// Algo type.
	Type string `json:"type"`

	// Used by Decode() method
	data []byte
}

func (model AlgoPhpass) New(data []byte) *AlgoPhpass {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AlgoPhpass) UnmarshalJSON(data []byte) error {
	type alias AlgoPhpass
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
		if err := json.Unmarshal([]byte("\"phpass\""), &decoded.Type); err != nil {
			return err
		}
	}

	*model = AlgoPhpass(decoded)
	return nil
}

func (model *AlgoPhpass) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
