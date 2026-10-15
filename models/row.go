package models

import (
	"encoding/json"
	"errors"
)

// Row Model
type Row struct {
	// Row ID.
	Id string `json:"$id"`
	// Row sequence ID.
	Sequence string `json:"$sequence"`
	// Table ID.
	TableId string `json:"$tableId"`
	// Database ID.
	DatabaseId string `json:"$databaseId"`
	// Row creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Row update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Row permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`

	// Used by Decode() method
	data []byte
}

func (model Row) New(data []byte) *Row {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Row) UnmarshalJSON(data []byte) error {
	type alias Row
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id         json.RawMessage `json:"$id"`
		Sequence   json.RawMessage `json:"$sequence"`
		TableId    json.RawMessage `json:"$tableId"`
		DatabaseId json.RawMessage `json:"$databaseId"`
		CreatedAt  json.RawMessage `json:"$createdAt"`
		UpdatedAt  json.RawMessage `json:"$updatedAt"`
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
	if present.TableId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TableId); err != nil {
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

	*model = Row(decoded)
	return nil
}

// Use this method to get response in desired type
func (model *Row) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
