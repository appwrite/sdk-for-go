package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Dropbox Model
type OAuth2Dropbox struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Dropbox OAuth2 app key.
	AppKey string `json:"appKey"`
	// Dropbox OAuth2 app secret.
	AppSecret string `json:"appSecret"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Dropbox) New(data []byte) *OAuth2Dropbox {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Dropbox) UnmarshalJSON(data []byte) error {
	type alias OAuth2Dropbox
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		Enabled   json.RawMessage `json:"enabled"`
		AppKey    json.RawMessage `json:"appKey"`
		AppSecret json.RawMessage `json:"appSecret"`
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
	if present.AppKey == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AppKey); err != nil {
			return err
		}
	}
	if present.AppSecret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AppSecret); err != nil {
			return err
		}
	}

	*model = OAuth2Dropbox(decoded)
	return nil
}

func (model *OAuth2Dropbox) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
