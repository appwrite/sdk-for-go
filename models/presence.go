package models

import (
	"encoding/json"
	"errors"
)

// Presence Model
type Presence struct {
	// Presence ID.
	Id string `json:"$id"`
	// Presence creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Presence update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Presence permissions. [Learn more about
	// permissions](https://appwrite.io/docs/permissions).
	Permissions []string `json:"$permissions"`
	// User ID.
	UserId string `json:"userId"`
	// Presence status.
	Status *string `json:"status"`
	// Presence source.
	Source string `json:"source"`
	// Presence expiry date in ISO 8601 format.
	ExpiresAt *string `json:"expiresAt"`
	// Presence metadata.
	Metadata interface{} `json:"metadata"`

	// Used by Decode() method
	data []byte
}

func (model Presence) New(data []byte) *Presence {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Presence) UnmarshalJSON(data []byte) error {
	type alias Presence
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UpdatedAt json.RawMessage `json:"$updatedAt"`
		UserId    json.RawMessage `json:"userId"`
		Source    json.RawMessage `json:"source"`
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
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.Source == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Source); err != nil {
			return err
		}
	}

	*model = Presence(decoded)
	return nil
}

func (model *Presence) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
