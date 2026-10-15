package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Gitlab Model
type OAuth2Gitlab struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// GitLab OAuth2 application ID.
	ApplicationId string `json:"applicationId"`
	// GitLab OAuth2 secret.
	Secret string `json:"secret"`
	// GitLab OAuth2 endpoint URL. Defaults to https://gitlab.com for self-hosted
	// instances.
	Endpoint string `json:"endpoint"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Gitlab) New(data []byte) *OAuth2Gitlab {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Gitlab) UnmarshalJSON(data []byte) error {
	type alias OAuth2Gitlab
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id            json.RawMessage `json:"$id"`
		Enabled       json.RawMessage `json:"enabled"`
		ApplicationId json.RawMessage `json:"applicationId"`
		Secret        json.RawMessage `json:"secret"`
		Endpoint      json.RawMessage `json:"endpoint"`
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
	if present.ApplicationId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ApplicationId); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}
	if present.Endpoint == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Endpoint); err != nil {
			return err
		}
	}

	*model = OAuth2Gitlab(decoded)
	return nil
}

func (model *OAuth2Gitlab) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
