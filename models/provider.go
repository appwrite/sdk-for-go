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

func (model *Provider) UnmarshalJSON(data []byte) error {
	type alias Provider
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$createdAt\":\"\",\"$updatedAt\":\"\",\"name\":\"\",\"provider\":\"\",\"enabled\":true,\"type\":\"\",\"credentials\":[]}"), &fields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	merged, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	return json.Unmarshal(merged, (*alias)(model))
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
