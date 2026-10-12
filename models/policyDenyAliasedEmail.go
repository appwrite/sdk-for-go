package models

import (
	"encoding/json"
	"errors"
)

// PolicyDenyAliasedEmail Model
type PolicyDenyAliasedEmail struct {
	// Policy ID.
	Id string `json:"$id"`
	// Whether the deny aliased email policy is enabled.
	Enabled bool `json:"enabled"`

	// Used by Decode() method
	data []byte
}

func (model PolicyDenyAliasedEmail) New(data []byte) *PolicyDenyAliasedEmail {
	model.data = data
	return &model
}

func (model *PolicyDenyAliasedEmail) UnmarshalJSON(data []byte) error {
	type alias PolicyDenyAliasedEmail
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false}"), &fields); err != nil {
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

func (model *PolicyDenyAliasedEmail) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
