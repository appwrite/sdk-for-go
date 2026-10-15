package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Notion Model
type OAuth2Notion struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Notion OAuth2 client ID.
	OauthClientId string `json:"oauthClientId"`
	// Notion OAuth2 client secret.
	OauthClientSecret string `json:"oauthClientSecret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Notion) New(data []byte) *OAuth2Notion {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Notion) UnmarshalJSON(data []byte) error {
	type alias OAuth2Notion
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                json.RawMessage `json:"$id"`
		Enabled           json.RawMessage `json:"enabled"`
		OauthClientId     json.RawMessage `json:"oauthClientId"`
		OauthClientSecret json.RawMessage `json:"oauthClientSecret"`
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
	if present.OauthClientId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.OauthClientId); err != nil {
			return err
		}
	}
	if present.OauthClientSecret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.OauthClientSecret); err != nil {
			return err
		}
	}

	*model = OAuth2Notion(decoded)
	return nil
}

func (model *OAuth2Notion) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
