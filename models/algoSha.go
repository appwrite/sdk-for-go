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

func (model *AlgoSha) UnmarshalJSON(data []byte) error {
	type alias AlgoSha
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"type\":\"sha\"}"), &fields); err != nil {
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
