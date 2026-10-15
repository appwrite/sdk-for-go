package models

import (
	"encoding/json"
	"errors"
)

// AlgoScrypt Model
type AlgoScrypt struct {
	// Algo type.
	Type string `json:"type"`
	// CPU complexity of computed hash.
	CostCpu int `json:"costCpu"`
	// Memory complexity of computed hash.
	CostMemory int `json:"costMemory"`
	// Parallelization of computed hash.
	CostParallel int `json:"costParallel"`
	// Length used to compute hash.
	Length int `json:"length"`

	// Used by Decode() method
	data []byte
}

func (model AlgoScrypt) New(data []byte) *AlgoScrypt {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AlgoScrypt) UnmarshalJSON(data []byte) error {
	type alias AlgoScrypt
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Type         json.RawMessage `json:"type"`
		CostCpu      json.RawMessage `json:"costCpu"`
		CostMemory   json.RawMessage `json:"costMemory"`
		CostParallel json.RawMessage `json:"costParallel"`
		Length       json.RawMessage `json:"length"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"scrypt\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.CostCpu == nil {
		if err := json.Unmarshal([]byte("8"), &decoded.CostCpu); err != nil {
			return err
		}
	}
	if present.CostMemory == nil {
		if err := json.Unmarshal([]byte("14"), &decoded.CostMemory); err != nil {
			return err
		}
	}
	if present.CostParallel == nil {
		if err := json.Unmarshal([]byte("1"), &decoded.CostParallel); err != nil {
			return err
		}
	}
	if present.Length == nil {
		if err := json.Unmarshal([]byte("64"), &decoded.Length); err != nil {
			return err
		}
	}

	*model = AlgoScrypt(decoded)
	return nil
}

func (model *AlgoScrypt) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
