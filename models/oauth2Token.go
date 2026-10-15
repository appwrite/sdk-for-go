package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Token Model
type Oauth2Token struct {
	// OAuth2 access token.
	AccessToken string `json:"access_token"`
	// OAuth2 token type.
	TokenType string `json:"token_type"`
	// Access token lifetime in seconds.
	ExpiresIn int `json:"expires_in"`
	// OAuth2 refresh token.
	RefreshToken string `json:"refresh_token"`
	// Space-separated scopes granted to the access token.
	Scope string `json:"scope"`
	// Granted RFC 9396 authorization details as a JSON string.
	AuthorizationDetails *string `json:"authorization_details"`
	// OpenID Connect ID token. Returned when the `openid` scope is granted.
	IdToken *string `json:"id_token"`

	// Used by Decode() method
	data []byte
}

func (model Oauth2Token) New(data []byte) *Oauth2Token {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Oauth2Token) UnmarshalJSON(data []byte) error {
	type alias Oauth2Token
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		AccessToken  json.RawMessage `json:"access_token"`
		TokenType    json.RawMessage `json:"token_type"`
		ExpiresIn    json.RawMessage `json:"expires_in"`
		RefreshToken json.RawMessage `json:"refresh_token"`
		Scope        json.RawMessage `json:"scope"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.AccessToken == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AccessToken); err != nil {
			return err
		}
	}
	if present.TokenType == nil {
		if err := json.Unmarshal([]byte("\"Bearer\""), &decoded.TokenType); err != nil {
			return err
		}
	}
	if present.ExpiresIn == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.ExpiresIn); err != nil {
			return err
		}
	}
	if present.RefreshToken == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RefreshToken); err != nil {
			return err
		}
	}
	if present.Scope == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Scope); err != nil {
			return err
		}
	}

	*model = Oauth2Token(decoded)
	return nil
}

func (model *Oauth2Token) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
