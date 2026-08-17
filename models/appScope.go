package models

import (
	"encoding/json"
	"errors"
)

// AppScope Model
type AppScope struct {
	// Scope value as requested by apps.
	Value string `json:"value"`
	// Human-readable description of what the scope grants.
	Description string `json:"description"`
	// What the scope grants access to. One of `account`, `project`, or
	// `organization`. Only `project` and `organization` scopes are installable.
	Type string `json:"type"`
	// Scope category, used to group scopes on consent and installation screens.
	Category string `json:"category"`
	// Whether the scope is deprecated. Deprecated scopes can still be requested
	// but should not be offered for new grants.
	Deprecated bool `json:"deprecated"`

	// Used by Decode() method
	data []byte
}

func (model AppScope) New(data []byte) *AppScope {
	model.data = data
	return &model
}

func (model *AppScope) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
