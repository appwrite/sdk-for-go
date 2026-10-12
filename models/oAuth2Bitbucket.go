package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Bitbucket Model
type OAuth2Bitbucket struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Bitbucket OAuth2 key.
	Key string `json:"key"`
	// Bitbucket OAuth2 secret.
	Secret string `json:"secret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Bitbucket) New(data []byte) *OAuth2Bitbucket {
	model.data = data
	return &model
}

func (model *OAuth2Bitbucket) UnmarshalJSON(data []byte) error {
	type alias OAuth2Bitbucket
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false,\"key\":\"\",\"secret\":\"\"}"), &fields); err != nil {
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

func (model *OAuth2Bitbucket) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
