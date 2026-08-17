package models

import (
	"encoding/json"
	"errors"
)

// PolicyDenyDisposableEmail Model
type PolicyDenyDisposableEmail struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether the deny disposable email policy is enabled.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model PolicyDenyDisposableEmail) New(data []byte) *PolicyDenyDisposableEmail {
	model.data = data
	return &model
}

func (model *PolicyDenyDisposableEmail) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
