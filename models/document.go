package models

import (
	"encoding/json"
	"errors"
)

// Document Model
type Document struct {
	// Document ID.
	Id string `json:"$id"`
	// Document sequence ID.
	Sequence string `json:"$sequence"`
	// Collection ID.
	CollectionId string `json:"$collectionId"`
	// Database ID.
	DatabaseId string `json:"$databaseId"`
	// Document creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Document update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Document permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`

	// Used by Decode() method
	data []byte
}

func (model Document) New(data []byte) *Document {
	model.data = data
	return &model
}

func (model *Document) UnmarshalJSON(data []byte) error {
	type alias Document
	var fields map[string]json.RawMessage
	if err := json.Unmarshal([]byte("{\"$id\":\"\",\"$sequence\":\"\",\"$collectionId\":\"\",\"$databaseId\":\"\",\"$createdAt\":\"\",\"$updatedAt\":\"\"}"), &fields); err != nil {
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

// Use this method to get response in desired type
func (model *Document) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
