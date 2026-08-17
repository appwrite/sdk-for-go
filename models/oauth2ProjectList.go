package models

import (
	"encoding/json"
	"errors"
)

// OAuth2AccessibleProjectsList Model
type Oauth2ProjectList struct {
	// Total number of projects that matched your query.
	Total int `json:"total"`
	// List of projects.
	Projects []Oauth2Project `json:"projects"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2ProjectList) New(data []byte) *Oauth2ProjectList {
	model.data = data
	return &model
}

func (model *Oauth2ProjectList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
