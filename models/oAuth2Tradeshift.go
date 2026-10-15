package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Tradeshift Model
type OAuth2Tradeshift struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Tradeshift OAuth2 client ID.
	Oauth2ClientId string `json:"oauth2ClientId"`
	// Tradeshift OAuth2 client secret.
	Oauth2ClientSecret string `json:"oauth2ClientSecret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Tradeshift) New(data []byte) *OAuth2Tradeshift {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Tradeshift) UnmarshalJSON(data []byte) error {
	type alias OAuth2Tradeshift
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                 json.RawMessage `json:"$id"`
		Enabled            json.RawMessage `json:"enabled"`
		Oauth2ClientId     json.RawMessage `json:"oauth2ClientId"`
		Oauth2ClientSecret json.RawMessage `json:"oauth2ClientSecret"`
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
	if present.Oauth2ClientId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Oauth2ClientId); err != nil {
			return err
		}
	}
	if present.Oauth2ClientSecret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Oauth2ClientSecret); err != nil {
			return err
		}
	}

	*model = OAuth2Tradeshift(decoded)
	return nil
}

func (model *OAuth2Tradeshift) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
