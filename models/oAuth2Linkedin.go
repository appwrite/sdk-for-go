package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Linkedin Model
type OAuth2Linkedin struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// LinkedIn OAuth2 client ID.
	ClientId string `json:"clientId"`
	// LinkedIn OAuth2 primary client secret.
	PrimaryClientSecret string `json:"primaryClientSecret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Linkedin) New(data []byte) *OAuth2Linkedin {
	model.data = data
	return &model
}

func (model *OAuth2Linkedin) UnmarshalJSON(data []byte) error {
	type alias OAuth2Linkedin
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false,\"clientId\":\"\",\"primaryClientSecret\":\"\"}"), &fields); err != nil {
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

func (model *OAuth2Linkedin) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
