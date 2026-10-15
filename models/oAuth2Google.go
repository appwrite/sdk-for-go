package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Google Model
type OAuth2Google struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Google OAuth2 client ID.
	ClientId string `json:"clientId"`
	// Google OAuth2 client secret.
	ClientSecret string `json:"clientSecret"`
	// Google OAuth2 prompt values.
	Prompt []string `json:"prompt"`
	// Native Google sign-in is active and can be used to create sessions from an
	// ID token. Independent of enabled, which only controls the browser-based
	// flow.
	NativeEnabled bool `json:"nativeEnabled"`
	// Additional OAuth2 client IDs accepted as ID token audiences for native
	// sign-in, next to the client ID.
	NativeClientIds []string `json:"nativeClientIds"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Google) New(data []byte) *OAuth2Google {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Google) UnmarshalJSON(data []byte) error {
	type alias OAuth2Google
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id              json.RawMessage `json:"$id"`
		Enabled         json.RawMessage `json:"enabled"`
		ClientId        json.RawMessage `json:"clientId"`
		ClientSecret    json.RawMessage `json:"clientSecret"`
		Prompt          json.RawMessage `json:"prompt"`
		NativeEnabled   json.RawMessage `json:"nativeEnabled"`
		NativeClientIds json.RawMessage `json:"nativeClientIds"`
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
		if err := json.Unmarshal([]byte("[\"consent\"]"), &decoded.Prompt); err != nil {
			return err
		}
	}
	if present.NativeEnabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.NativeEnabled); err != nil {
			return err
		}
	}
	if present.NativeClientIds == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.NativeClientIds); err != nil {
			return err
		}
	}

	*model = OAuth2Google(decoded)
	return nil
}

func (model *OAuth2Google) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
