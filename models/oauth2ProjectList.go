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

func (model *Oauth2ProjectList) UnmarshalJSON(data []byte) error {
	type alias Oauth2ProjectList
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"total\":0}"), &fields); err != nil {
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
