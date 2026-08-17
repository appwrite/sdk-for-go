package models

import (
	"encoding/json"
	"errors"
)

// OAuth2PAR Model
type Oauth2PAR struct {
	// Authorization request handle to pass to the authorize endpoint.
	RequestUri string `json:"request_uri"`
	// Lifetime of the authorization request handle in seconds.
	ExpiresIn int `json:"expires_in"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2PAR) New(data []byte) *Oauth2PAR {
	model.data = data
	return &model
}

func (model *Oauth2PAR) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
