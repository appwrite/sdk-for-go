package models

import (
	"encoding/json"
	"errors"
)

// AlgoMD5 Model
type AlgoMd5 struct {
	// Algo type.
	Type string `json:"type"`

	// Used by Decode() method
	data []byte
}

func (model AlgoMd5) New(data []byte) *AlgoMd5 {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AlgoMd5) UnmarshalJSON(data []byte) error {
	type alias AlgoMd5
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
		if err := json.Unmarshal([]byte("\"md5\""), &decoded.Type); err != nil {
			return err
		}
	}

	*model = AlgoMd5(decoded)
	return nil
}

func (model *AlgoMd5) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
