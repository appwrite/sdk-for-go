package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Okta Model
type OAuth2Okta struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Okta OAuth2 client ID.
	ClientId string `json:"clientId"`
	// Okta OAuth2 client secret.
	ClientSecret string `json:"clientSecret"`
	// Okta OAuth2 prompt values.
	Prompt []string `json:"prompt"`
	// Okta OAuth2 domain.
	Domain string `json:"domain"`
	// Okta OAuth2 authorization server ID.
	AuthorizationServerId string `json:"authorizationServerId"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Okta) New(data []byte) *OAuth2Okta {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Okta) UnmarshalJSON(data []byte) error {
	type alias OAuth2Okta
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                    json.RawMessage `json:"$id"`
		Enabled               json.RawMessage `json:"enabled"`
		ClientId              json.RawMessage `json:"clientId"`
		ClientSecret          json.RawMessage `json:"clientSecret"`
		Prompt                json.RawMessage `json:"prompt"`
		Domain                json.RawMessage `json:"domain"`
		AuthorizationServerId json.RawMessage `json:"authorizationServerId"`
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
	if present.Domain == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Domain); err != nil {
			return err
		}
	}
	if present.AuthorizationServerId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AuthorizationServerId); err != nil {
			return err
		}
	}

	*model = OAuth2Okta(decoded)
	return nil
}

func (model *OAuth2Okta) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
