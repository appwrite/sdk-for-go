package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Keycloak Model
type OAuth2Keycloak struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Keycloak OAuth2 client ID.
	ClientId string `json:"clientId"`
	// Keycloak OAuth2 client secret.
	ClientSecret string `json:"clientSecret"`
	// Keycloak OAuth2 endpoint domain.
	Endpoint string `json:"endpoint"`
	// Keycloak OAuth2 realm name.
	RealmName string `json:"realmName"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Keycloak) New(data []byte) *OAuth2Keycloak {
	model.data = data
	return &model
}

func (model *OAuth2Keycloak) UnmarshalJSON(data []byte) error {
	type alias OAuth2Keycloak
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false,\"clientId\":\"\",\"clientSecret\":\"\",\"endpoint\":\"\",\"realmName\":\"\"}"), &fields); err != nil {
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

func (model *OAuth2Keycloak) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
