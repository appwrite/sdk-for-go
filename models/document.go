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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Document) UnmarshalJSON(data []byte) error {
	type alias Document
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		Sequence     json.RawMessage `json:"$sequence"`
		CollectionId json.RawMessage `json:"$collectionId"`
		DatabaseId   json.RawMessage `json:"$databaseId"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.Sequence == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Sequence); err != nil {
			return err
		}
	}
	if present.CollectionId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CollectionId); err != nil {
			return err
		}
	}
	if present.DatabaseId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DatabaseId); err != nil {
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

	*model = Document(decoded)
	return nil
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
