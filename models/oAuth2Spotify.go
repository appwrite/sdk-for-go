package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Spotify Model
type OAuth2Spotify struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Spotify OAuth2 client ID.
	ClientId string `json:"clientId"`
	// Spotify OAuth2 client secret.
	ClientSecret string `json:"clientSecret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Spotify) New(data []byte) *OAuth2Spotify {
	model.data = data
	return &model
}

func (model *OAuth2Spotify) UnmarshalJSON(data []byte) error {
	type alias OAuth2Spotify
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false,\"clientId\":\"\",\"clientSecret\":\"\"}"), &fields); err != nil {
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

func (model *OAuth2Spotify) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
