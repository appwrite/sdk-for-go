package models

import (
	"encoding/json"
	"errors"
)

// OAuth2TikTok Model
type OAuth2TikTok struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// TikTok OAuth2 client key.
	ClientId string `json:"clientId"`
	// TikTok OAuth2 client secret.
	ClientSecret string `json:"clientSecret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2TikTok) New(data []byte) *OAuth2TikTok {
	model.data = data
	return &model
}

func (model *OAuth2TikTok) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
