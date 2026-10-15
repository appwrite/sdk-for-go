package models

import (
	"encoding/json"
	"errors"
)

// Provider Model
type Provider struct {
	// Provider ID.
	Id string `json:"$id"`
	// Provider creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Provider update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// The name for the provider instance.
	Name string `json:"name"`
	// The name of the provider service.
	Provider string `json:"provider"`
	// Is provider enabled?
	Enabled bool `json:"enabled"`
	// Type of provider.
	Type string `json:"type"`
	// Provider credentials.
	Credentials interface{} `json:"credentials"`
	// Provider options.
	Options interface{} `json:"options"`

	// Used by Decode() method
	data []byte
}

func (model Provider) New(data []byte) *Provider {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Provider) UnmarshalJSON(data []byte) error {
	type alias Provider
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id          json.RawMessage `json:"$id"`
		CreatedAt   json.RawMessage `json:"$createdAt"`
		UpdatedAt   json.RawMessage `json:"$updatedAt"`
		Name        json.RawMessage `json:"name"`
		Provider    json.RawMessage `json:"provider"`
		Enabled     json.RawMessage `json:"enabled"`
		Type        json.RawMessage `json:"type"`
		Credentials json.RawMessage `json:"credentials"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Provider == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Provider); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Credentials == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Credentials); err != nil {
			return err
		}
	}

	*model = Provider(decoded)
	return nil
}

func (model *Provider) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
