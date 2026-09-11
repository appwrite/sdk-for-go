package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Kakao Model
type OAuth2Kakao struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Kakao OAuth2 REST API key.
	ClientId string `json:"clientId"`
	// Kakao OAuth2 client secret.
	ClientSecret string `json:"clientSecret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Kakao) New(data []byte) *OAuth2Kakao {
	model.data = data
	return &model
}

func (model *OAuth2Kakao) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
