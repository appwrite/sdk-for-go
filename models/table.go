package models

import (
	"encoding/json"
	"errors"
)

// Table Model
type Table struct {
	// Table ID.
	Id string `json:"$id"`
	// Table creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Table update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Table permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`
	// Database ID.
	DatabaseId string `json:"databaseId"`
	// Table name.
	Name string `json:"name"`
	// Optional description for the table.
	Description string `json:"description"`
	// Table enabled. Can be 'enabled' or 'disabled'. When disabled, the table is
	// inaccessible to users, but remains accessible to Server SDKs using API
	// keys.
	Enabled bool `json:"enabled"`
	// Whether row-level permissions are enabled. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	RowSecurity bool `json:"rowSecurity"`
	// Table columns.
	Columns []interface{} `json:"columns"`
	// Table indexes.
	Indexes []ColumnIndex `json:"indexes"`
	// Maximum row size in bytes. Returns 0 when no limit applies.
	BytesMax int `json:"bytesMax"`
	// Currently used row size in bytes based on defined columns.
	BytesUsed int `json:"bytesUsed"`

	// Used by Decode() method
	data []byte
}

func (model Table) New(data []byte) *Table {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Table) UnmarshalJSON(data []byte) error {
	type alias Table
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id          json.RawMessage `json:"$id"`
		CreatedAt   json.RawMessage `json:"$createdAt"`
		UpdatedAt   json.RawMessage `json:"$updatedAt"`
		DatabaseId  json.RawMessage `json:"databaseId"`
		Name        json.RawMessage `json:"name"`
		Description json.RawMessage `json:"description"`
		Enabled     json.RawMessage `json:"enabled"`
		RowSecurity json.RawMessage `json:"rowSecurity"`
		BytesMax    json.RawMessage `json:"bytesMax"`
		BytesUsed   json.RawMessage `json:"bytesUsed"`
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
	if present.Description == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Description); err != nil {
			return err
		}
	}
	if present.Enabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Enabled); err != nil {
			return err
		}
	}
	if present.RowSecurity == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.RowSecurity); err != nil {
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

	*model = Table(decoded)
	return nil
}

func (model *Table) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
