package models

import (
	"encoding/json"
	"errors"
)

// Passkey Model
type Passkey struct {
	// Passkey ID.
	Id string `json:"$id"`
	// Passkey creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Passkey update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Passkey name.
	Name string `json:"name"`
	// Most recent sign-in with this passkey in ISO 8601 format. Empty until the
	// passkey is used to sign in.
	AccessedAt string `json:"accessedAt"`

	// Used by Decode() method
	data []byte
}

func (model Passkey) New(data []byte) *Passkey {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Passkey) UnmarshalJSON(data []byte) error {
	type alias Passkey
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id         json.RawMessage `json:"$id"`
		CreatedAt  json.RawMessage `json:"$createdAt"`
		UpdatedAt  json.RawMessage `json:"$updatedAt"`
		Name       json.RawMessage `json:"name"`
		AccessedAt json.RawMessage `json:"accessedAt"`
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
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.AccessedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AccessedAt); err != nil {
			return err
		}
	}

	*model = Passkey(decoded)
	return nil
}

func (model *Passkey) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
