package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Facebook Model
type OAuth2Facebook struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Facebook OAuth2 app ID.
	AppId string `json:"appId"`
	// Facebook OAuth2 app secret.
	AppSecret string `json:"appSecret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Facebook) New(data []byte) *OAuth2Facebook {
	model.data = data
	return &model
}

func (model *OAuth2Facebook) UnmarshalJSON(data []byte) error {
	type alias OAuth2Facebook
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"enabled\":false,\"appId\":\"\",\"appSecret\":\"\"}"), &fields); err != nil {
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

func (model *OAuth2Facebook) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
