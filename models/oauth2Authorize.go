package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Authorize Model
type Oauth2Authorize struct {
	// OAuth2 grant ID. Set when the user must give explicit consent; pass it to
	// the approve or reject endpoint. Empty when a redirect URL is returned
	// instead.
	GrantId string `json:"grantId"`
	// URL the end user should be redirected to when the flow can complete without
	// consent. Empty when consent is still required.
	RedirectUrl string `json:"redirectUrl"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Authorize) New(data []byte) *Oauth2Authorize {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2Authorize) UnmarshalJSON(data []byte) error {
	type alias Oauth2Authorize
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		GrantId     json.RawMessage `json:"grantId"`
		RedirectUrl json.RawMessage `json:"redirectUrl"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.GrantId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.GrantId); err != nil {
			return err
		}
	}
	if present.RedirectUrl == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RedirectUrl); err != nil {
			return err
		}
	}

	*model = Oauth2Authorize(decoded)
	return nil
}

func (model *Oauth2Authorize) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
