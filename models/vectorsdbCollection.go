package models

import (
	"encoding/json"
	"errors"
)

// VectorsDBCollection Model
type VectorsdbCollection struct {
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
	// Embedding dimension.
	Dimension int `json:"dimension"`

	// Used by Decode() method
	data []byte
}

func (model VectorsdbCollection) New(data []byte) *VectorsdbCollection {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *VectorsdbCollection) UnmarshalJSON(data []byte) error {
	type alias VectorsdbCollection
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id               json.RawMessage `json:"$id"`
		CreatedAt        json.RawMessage `json:"$createdAt"`
		UpdatedAt        json.RawMessage `json:"$updatedAt"`
		DatabaseId       json.RawMessage `json:"databaseId"`
		Name             json.RawMessage `json:"name"`
		Enabled          json.RawMessage `json:"enabled"`
		DocumentSecurity json.RawMessage `json:"documentSecurity"`
		BytesMax         json.RawMessage `json:"bytesMax"`
		BytesUsed        json.RawMessage `json:"bytesUsed"`
		Dimension        json.RawMessage `json:"dimension"`
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
	if present.DatabaseId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.DatabaseId); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.DocumentSecurity == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.DocumentSecurity); err != nil {
			return err
		}
	}
	if present.BytesMax == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.BytesMax); err != nil {
			return err
		}
	}
	if present.BytesUsed == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.BytesUsed); err != nil {
			return err
		}
	}
	if present.Dimension == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Dimension); err != nil {
			return err
		}
	}

	*model = VectorsdbCollection(decoded)
	return nil
}

func (model *VectorsdbCollection) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
