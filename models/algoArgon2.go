package models

import (
	"encoding/json"
	"errors"
)

// AlgoArgon2 Model
type AlgoArgon2 struct {
	// Algo type.
	Type string `json:"type"`
	// Memory used to compute hash.
	MemoryCost int `json:"memoryCost"`
	// Amount of time consumed to compute hash
	TimeCost int `json:"timeCost"`
	// Number of threads used to compute hash.
	Threads int `json:"threads"`

	// Used by Decode() method
	data []byte
}

func (model AlgoArgon2) New(data []byte) *AlgoArgon2 {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AlgoArgon2) UnmarshalJSON(data []byte) error {
	type alias AlgoArgon2
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Type       json.RawMessage `json:"type"`
		MemoryCost json.RawMessage `json:"memoryCost"`
		TimeCost   json.RawMessage `json:"timeCost"`
		Threads    json.RawMessage `json:"threads"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"argon2\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.MemoryCost == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.MemoryCost); err != nil {
			return err
		}
	}
	if present.TimeCost == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.TimeCost); err != nil {
			return err
		}
	}
	if present.Threads == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Threads); err != nil {
			return err
		}
	}

	*model = AlgoArgon2(decoded)
	return nil
}

func (model *AlgoArgon2) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
