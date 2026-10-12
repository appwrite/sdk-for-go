package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Project Model
type Oauth2Project struct {
	// Project ID.
	Id string `json:"$id"`
	// Region ID the project is deployed in.
	Region string `json:"region"`
	// API endpoint of the region the project is deployed in. Empty when the
	// region has no public hostname configured.
	Endpoint string `json:"endpoint"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Project) New(data []byte) *Oauth2Project {
	model.data = data
	return &model
}

func (model *Oauth2Project) UnmarshalJSON(data []byte) error {
	type alias Oauth2Project
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"region\":\"\",\"endpoint\":\"\"}"), &fields); err != nil {
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

func (model *Oauth2Project) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
