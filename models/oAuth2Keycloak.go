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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Keycloak) UnmarshalJSON(data []byte) error {
	type alias OAuth2Keycloak
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		Enabled      json.RawMessage `json:"enabled"`
		ClientId     json.RawMessage `json:"clientId"`
		ClientSecret json.RawMessage `json:"clientSecret"`
		Endpoint     json.RawMessage `json:"endpoint"`
		RealmName    json.RawMessage `json:"realmName"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.ClientId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientId); err != nil {
			return err
		}
	}
	if present.ClientSecret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientSecret); err != nil {
			return err
		}
	}
	if present.Endpoint == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Endpoint); err != nil {
			return err
		}
	}
	if present.RealmName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RealmName); err != nil {
			return err
		}
	}

	*model = OAuth2Keycloak(decoded)
	return nil
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
