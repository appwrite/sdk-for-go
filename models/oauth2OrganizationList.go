package models

import (
	"encoding/json"
	"errors"
)

// OAuth2AccessibleOrganizationsList Model
type Oauth2OrganizationList struct {
	// Total number of organizations that matched your query.
	Total int `json:"total"`
	// List of organizations.
	Organizations []Oauth2Organization `json:"organizations"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2OrganizationList) New(data []byte) *Oauth2OrganizationList {
	model.data = data
	return &model
}

func (model *Oauth2OrganizationList) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
