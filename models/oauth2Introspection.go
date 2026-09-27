package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Introspection Model
type Oauth2Introspection struct {
	// Whether the token is currently active. An inactive token carries no other
	// claims (RFC 7662).
	Active bool `json:"active"`
	// Kind of token that was introspected. Can be one of: `access_token` or
	// `refresh_token`.
	TokenUse *string `json:"token_use"`
	// OAuth2 token type.
	TokenType *string `json:"token_type"`
	// Space-separated scopes granted to the token.
	Scope *string `json:"scope"`
	// OAuth2 client ID the token was issued to.
	ClientId *string `json:"client_id"`
	// ID of the user who authorized the token.
	Sub *string `json:"sub"`
	// Audiences the token is intended for. The first entry is the project API
	// base URL; any further entries are the resource indicators requested at
	// authorization time (RFC 8707).
	Aud []string `json:"aud"`
	// Issuer URL of the token.
	Iss *string `json:"iss"`
	// Expiration time as a Unix timestamp in seconds.
	Exp *int `json:"exp"`
	// Issued-at time as a Unix timestamp in seconds.
	Iat *int `json:"iat"`
	// Unique identifier of the token.
	Jti *string `json:"jti"`
	// Granted RFC 9396 authorization details, restricted to what the user can
	// currently access.
	AuthorizationDetails []interface{} `json:"authorization_details"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Introspection) New(data []byte) *Oauth2Introspection {
	model.data = data
	return &model
}

func (model *Oauth2Introspection) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
