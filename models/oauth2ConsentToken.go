package models

import (
	"encoding/json"
	"errors"
)

// OAuth2ConsentToken Model
type Oauth2ConsentToken struct {
	// Token family ID.
	Id string `json:"$id"`
	// Token creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Token update date in ISO 8601 format. Refreshing the token family updates
	// this.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the consent the token family was issued under.
	ConsentId string `json:"consentId"`
	// ID of the user the token family belongs to.
	UserId string `json:"userId"`
	// ID of the registered app the token family was issued to. Empty for URL-form
	// (CIMD) clients.
	AppId string `json:"appId"`
	// Client ID metadata document URL of the client the token family was issued
	// to. Empty for registered apps.
	CimdUrl string `json:"cimdUrl"`
	// OAuth2 scopes granted on the token family.
	Scopes []string `json:"scopes"`
	// RFC 8707 resource indicators granted on the token family.
	Resources []string `json:"resources"`
	// Authorization details granted on the token family, as a JSON string. Each
	// entry has a `type` plus project-defined fields.
	AuthorizationDetails string `json:"authorizationDetails"`
	// Expiration time of the current access token of this family in ISO 8601
	// format.
	Expire string `json:"expire"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2ConsentToken) New(data []byte) *Oauth2ConsentToken {
	model.data = data
	return &model
}

func (model *Oauth2ConsentToken) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
