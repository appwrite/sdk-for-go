package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Apple Model
type OAuth2Apple struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Apple OAuth2 service ID.
	ServiceId string `json:"serviceId"`
	// Apple OAuth2 key ID.
	KeyId string `json:"keyId"`
	// Apple OAuth2 team ID.
	TeamId string `json:"teamId"`
	// Apple OAuth2 .p8 private key file contents. The secret key wrapped by the
	// PEM markers is 200 characters long.
	P8File string `json:"p8File"`
	// Native Sign in with Apple is active and can be used to create sessions from
	// an ID token. Independent of enabled, which only controls the browser-based
	// flow.
	NativeEnabled bool `json:"nativeEnabled"`
	// App bundle IDs accepted as ID token audiences for native Sign in with
	// Apple, next to the Services ID.
	NativeClientIds []string `json:"nativeClientIds"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Apple) New(data []byte) *OAuth2Apple {
	model.data = data
	return &model
}

func (model *OAuth2Apple) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
