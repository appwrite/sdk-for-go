package models

import (
	"encoding/json"
	"errors"
)

// AppKey Model
type AppKey struct {
	// App key ID.
	Id string `json:"$id"`
	// App key creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// App key update time in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Application ID this app key belongs to.
	AppId string `json:"appId"`
	// App key secret.
	Secret string `json:"secret"`
	// Last few characters of the app key secret, used to help identify it.
	Hint string `json:"hint"`
	// ID of the user who created the app key.
	CreatedById string `json:"createdById"`
	// Name of the user who created the app key.
	CreatedByName string `json:"createdByName"`
	// Time the app key was last used for authentication in ISO 8601 format. Null
	// if never used.
	LastAccessedAt string `json:"lastAccessedAt"`

	// Used by Decode() method
	data []byte
}

func (model AppKey) New(data []byte) *AppKey {
	model.data = data
	return &model
}

func (model *AppKey) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
