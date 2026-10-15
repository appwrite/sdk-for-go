package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Oidc Model
type OAuth2Oidc struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// OpenID Connect OAuth2 client ID.
	ClientId string `json:"clientId"`
	// OpenID Connect OAuth2 client secret.
	ClientSecret string `json:"clientSecret"`
	// OpenID Connect prompt values controlling the authentication and consent
	// screens.
	Prompt []string `json:"prompt"`
	// OpenID Connect well-known configuration URL. When set, authorization,
	// token, and user info endpoints can be discovered automatically.
	WellKnownURL string `json:"wellKnownURL"`
	// OpenID Connect authorization endpoint URL.
	AuthorizationURL string `json:"authorizationURL"`
	// OpenID Connect token endpoint URL.
	TokenURL string `json:"tokenURL"`
	// OpenID Connect user info endpoint URL.
	UserInfoURL string `json:"userInfoURL"`
	// Maximum authentication age in seconds. When set, the user must have
	// authenticated within this many seconds.
	MaxAge *int `json:"maxAge"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Oidc) New(data []byte) *OAuth2Oidc {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Oidc) UnmarshalJSON(data []byte) error {
	type alias OAuth2Oidc
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id               json.RawMessage `json:"$id"`
		Enabled          json.RawMessage `json:"enabled"`
		ClientId         json.RawMessage `json:"clientId"`
		ClientSecret     json.RawMessage `json:"clientSecret"`
		Prompt           json.RawMessage `json:"prompt"`
		WellKnownURL     json.RawMessage `json:"wellKnownURL"`
		AuthorizationURL json.RawMessage `json:"authorizationURL"`
		TokenURL         json.RawMessage `json:"tokenURL"`
		UserInfoURL      json.RawMessage `json:"userInfoURL"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.ClientId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientId); err != nil {
			return err
		}
	}
	if present.ClientSecret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ClientSecret); err != nil {
			return err
		}
	}
	if present.Prompt == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Prompt); err != nil {
			return err
		}
	}
	if present.WellKnownURL == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.WellKnownURL); err != nil {
			return err
		}
	}
	if present.AuthorizationURL == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AuthorizationURL); err != nil {
			return err
		}
	}
	if present.TokenURL == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TokenURL); err != nil {
			return err
		}
	}
	if present.UserInfoURL == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserInfoURL); err != nil {
			return err
		}
	}

	*model = OAuth2Oidc(decoded)
	return nil
}

func (model *OAuth2Oidc) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
