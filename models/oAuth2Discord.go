package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Discord Model
type OAuth2Discord struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Discord OAuth2 client ID.
	ClientId string `json:"clientId"`
	// Discord OAuth2 client secret.
	ClientSecret string `json:"clientSecret"`
	// Discord OAuth2 prompt values.
	Prompt []string `json:"prompt"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Discord) New(data []byte) *OAuth2Discord {
	model.data = data
	return &model
}

func (model *OAuth2Discord) UnmarshalJSON(data []byte) error {
	type alias OAuth2Discord
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false,\"clientId\":\"\",\"clientSecret\":\"\",\"prompt\":[]}"), &fields); err != nil {
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

func (model *OAuth2Discord) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
