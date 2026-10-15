package models

import (
	"encoding/json"
	"errors"
)

// Key Model
type Key struct {
	// Key ID.
	Id string `json:"$id"`
	// Key creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Key update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Key name.
	Name string `json:"name"`
	// Key expiration date in ISO 8601 format.
	Expire string `json:"expire"`
	// Allowed permission scopes.
	Scopes []string `json:"scopes"`
	// Secret key.
	Secret string `json:"secret"`
	// Most recent access date in ISO 8601 format. This attribute is only updated
	// again after 24 hours.
	AccessedAt string `json:"accessedAt"`
	// List of SDK user agents that used this key.
	Sdks []string `json:"sdks"`

	// Used by Decode() method
	data []byte
}

func (model Key) New(data []byte) *Key {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Key) UnmarshalJSON(data []byte) error {
	type alias Key
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id         json.RawMessage `json:"$id"`
		CreatedAt  json.RawMessage `json:"$createdAt"`
		UpdatedAt  json.RawMessage `json:"$updatedAt"`
		Name       json.RawMessage `json:"name"`
		Expire     json.RawMessage `json:"expire"`
		Scopes     json.RawMessage `json:"scopes"`
		Secret     json.RawMessage `json:"secret"`
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
	if present.Expire == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Expire); err != nil {
			return err
		}
	}
	if present.Scopes == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.Scopes); err != nil {
			return err
		}
	}
	if present.Secret == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Secret); err != nil {
			return err
		}
	}
	if present.AccessedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.AccessedAt); err != nil {
			return err
		}
	}

	*model = Key(decoded)
	return nil
}

func (model *Key) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
