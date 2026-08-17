package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Organization Model
type Oauth2Organization struct {
	// Organization ID.
	Id string `json:"$id"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Organization) New(data []byte) *Oauth2Organization {
	model.data = data
	return &model
}

func (model *Oauth2Organization) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
