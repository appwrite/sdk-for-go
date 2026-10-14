package models

import (
	"encoding/json"
	"errors"
)

// Collection Model
type Collection struct {
	// Collection ID.
	Id string `json:"$id"`
	// Collection creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Collection update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Collection permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`
	// Database ID.
	DatabaseId string `json:"databaseId"`
	// Collection name.
	Name string `json:"name"`
	// Collection enabled. Can be 'enabled' or 'disabled'. When disabled, the
	// collection is inaccessible to users, but remains accessible to Server SDKs
	// using API keys.
	Enabled bool `json:"enabled"`
	// Whether document-level permissions are enabled. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	DocumentSecurity bool `json:"documentSecurity"`
	// Collection attributes.
	Attributes []map[string]any `json:"attributes"`
	// Collection indexes.
	Indexes []Index `json:"indexes"`
	// Maximum document size in bytes. Returns 0 when no limit applies.
	BytesMax int `json:"bytesMax"`
	// Currently used document size in bytes based on defined attributes.
	BytesUsed int `json:"bytesUsed"`

	// Used by Decode() method
	data []byte
}

func (model Collection) New(data []byte) *Collection {
	model.data = data
	return &model
}

func (model *Collection) UnmarshalJSON(data []byte) error {
	type alias Collection
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$createdAt\":\"\",\"$updatedAt\":\"\",\"databaseId\":\"\",\"name\":\"\",\"enabled\":true,\"documentSecurity\":false,\"bytesMax\":0,\"bytesUsed\":0}"), &fields); err != nil {
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

func (model *Collection) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
