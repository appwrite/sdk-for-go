package models

import (
	"encoding/json"
	"errors"
)

// OAuth2ConsentsList Model
type Oauth2ConsentList struct {
	// Total number of consents that matched your query.
	Total int `json:"total"`
	// List of consents.
	Consents []Oauth2Consent `json:"consents"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2ConsentList) New(data []byte) *Oauth2ConsentList {
	model.data = data
	return &model
}

func (model *Oauth2ConsentList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
