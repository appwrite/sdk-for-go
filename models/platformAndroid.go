package models

import (
	"encoding/json"
	"errors"
)

// PlatformAndroid Model
type PlatformAndroid struct {
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
	// Android application ID.
	ApplicationId string `json:"applicationId"`
	// SHA-256 fingerprints of the certificates the app is signed with. Passkeys
	// created in the app are accepted when its signature matches one.
	Fingerprints []string `json:"fingerprints"`

	// Used by Decode() method
	data []byte
}

func (model PlatformAndroid) New(data []byte) *PlatformAndroid {
	model.data = data
	return &model
}

func (model *PlatformAndroid) UnmarshalJSON(data []byte) error {
	type alias PlatformAndroid
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$createdAt\":\"\",\"$updatedAt\":\"\",\"name\":\"\",\"applicationId\":\"\",\"fingerprints\":[]}"), &fields); err != nil {
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

func (model *PlatformAndroid) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
