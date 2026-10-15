package models

import (
	"encoding/json"
	"errors"
)

// OAuth2Microsoft Model
type OAuth2Microsoft struct {
	// OAuth2 provider ID.
	Id string `json:"$id"`
	// OAuth2 provider is active and can be used to create sessions.
	Enabled bool `json:"enabled"`
	// Microsoft OAuth2 application ID.
	ApplicationId string `json:"applicationId"`
	// Microsoft OAuth2 application secret.
	ApplicationSecret string `json:"applicationSecret"`
	// Microsoft OAuth2 prompt values.
	Prompt []string `json:"prompt"`
	// Microsoft Entra ID tenant identifier. Use 'common', 'organizations',
	// 'consumers' or a specific tenant ID.
	Tenant string `json:"tenant"`

	// Used by Decode() method
	data []byte
}

func (model OAuth2Microsoft) New(data []byte) *OAuth2Microsoft {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *OAuth2Microsoft) UnmarshalJSON(data []byte) error {
	type alias OAuth2Microsoft
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                json.RawMessage `json:"$id"`
		Enabled           json.RawMessage `json:"enabled"`
		ApplicationId     json.RawMessage `json:"applicationId"`
		ApplicationSecret json.RawMessage `json:"applicationSecret"`
		Prompt            json.RawMessage `json:"prompt"`
		Tenant            json.RawMessage `json:"tenant"`
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
	if present.ApplicationSecret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ApplicationSecret); err != nil {
			return err
		}
	}
	if present.Prompt == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Prompt); err != nil {
			return err
		}
	}
	if present.Tenant == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Tenant); err != nil {
			return err
		}
	}

	*model = OAuth2Microsoft(decoded)
	return nil
}

func (model *OAuth2Microsoft) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
