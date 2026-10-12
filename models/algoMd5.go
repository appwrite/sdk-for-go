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

func (model *AlgoMd5) UnmarshalJSON(data []byte) error {
	type alias AlgoMd5
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"type\":\"md5\"}"), &fields); err != nil {
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
