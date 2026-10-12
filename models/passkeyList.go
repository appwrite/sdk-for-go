package models

import (
	"encoding/json"
	"errors"
)

// PasskeysList Model
type PasskeyList struct {
	// Total number of passkeys that matched your query.
	Total int `json:"total"`
	// List of passkeys.
	Passkeys []Passkey `json:"passkeys"`

	// Used by Decode() method
	data []byte
}

func (model PasskeyList) New(data []byte) *PasskeyList {
	model.data = data
	return &model
}

func (model *PasskeyList) UnmarshalJSON(data []byte) error {
	type alias PasskeyList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
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

func (model *PasskeyList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
