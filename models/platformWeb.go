package models

import (
	"encoding/json"
	"errors"
)

// PlatformWeb Model
type PlatformWeb struct {
	// Platform ID.
	Id string `json:"$id"`
	// Platform creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Platform update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Platform name.
	Name string `json:"name"`
	// Platform type. Possible values are: windows, apple, android, linux, web.
	Type string `json:"type"`
	// Web app hostname. Empty string for other platforms.
	Hostname string `json:"hostname"`

	// Used by Decode() method
	data []byte
}

func (model PlatformWeb) New(data []byte) *PlatformWeb {
	model.data = data
	return &model
}

func (model *PlatformWeb) UnmarshalJSON(data []byte) error {
	type alias PlatformWeb
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$createdAt\":\"\",\"$updatedAt\":\"\",\"name\":\"\",\"hostname\":\"\"}"), &fields); err != nil {
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

func (model *PlatformWeb) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
