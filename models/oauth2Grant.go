package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Grant Model
type Oauth2Grant struct {
	// Grant ID.
	Id string `json:"$id"`
	// Grant creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Grant update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the user the grant belongs to.
	UserId string `json:"userId"`
	// ID of the OAuth2 client (app) the grant was requested for.
	AppId string `json:"appId"`
	// Requested OAuth2 scopes the user is being asked to consent to.
	Scopes []string `json:"scopes"`
	// Requested RFC 8707 resource indicators the user is being asked to consent
	// to.
	Resources []string `json:"resources"`
	// Requested authorization_details the user is being asked to consent to, as a
	// JSON string. Each entry has a `type` plus project-defined fields.
	AuthorizationDetails string `json:"authorizationDetails"`
	// OIDC prompt directive the consent screen should honor. Space-separated list
	// of: login, consent, select_account.
	Prompt string `json:"prompt"`
	// Redirect URI the user will be sent to after the flow completes.
	RedirectUri string `json:"redirectUri"`
	// Unix timestamp of when the user last authenticated.
	AuthTime int `json:"authTime"`
	// Grant expiration time in ISO 8601 format.
	Expire string `json:"expire"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Grant) New(data []byte) *Oauth2Grant {
	model.data = data
	return &model
}

func (model *Oauth2Grant) UnmarshalJSON(data []byte) error {
	type alias Oauth2Grant
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$createdAt\":\"\",\"$updatedAt\":\"\",\"userId\":\"\",\"appId\":\"\",\"scopes\":[],\"resources\":[],\"authorizationDetails\":\"[]\",\"prompt\":\"\",\"redirectUri\":\"\",\"authTime\":0,\"expire\":\"\"}"), &fields); err != nil {
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

func (model *Oauth2Grant) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
