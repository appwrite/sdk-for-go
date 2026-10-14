package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Approve Model
type Oauth2Approve struct {
	// URL the end user should be redirected to after the grant is approved,
	// carrying the authorization `code` and/or `id_token` along with the original
	// `state`.
	RedirectUrl string `json:"redirectUrl"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Approve) New(data []byte) *Oauth2Approve {
	model.data = data
	return &model
}

func (model *Oauth2Approve) UnmarshalJSON(data []byte) error {
	type alias Oauth2Approve
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"redirectUrl\":\"\"}"), &fields); err != nil {
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

func (model *Oauth2Approve) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
