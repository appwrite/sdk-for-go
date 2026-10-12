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

func (model *Oauth2Organization) UnmarshalJSON(data []byte) error {
	type alias Oauth2Organization
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\"}"), &fields); err != nil {
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
