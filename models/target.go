package models

import (
	"encoding/json"
	"errors"
)

// Target Model
type Target struct {
	// Target ID.
	Id string `json:"$id"`
	// Target creation time in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Target update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Target Name.
	Name string `json:"name"`
	// User ID.
	UserId string `json:"userId"`
	// Provider ID.
	ProviderId *string `json:"providerId"`
	// The target provider type. Can be one of the following: `email`, `sms` or
	// `push`.
	ProviderType string `json:"providerType"`
	// The target identifier.
	Identifier string `json:"identifier"`
	// Is the target expired.
	Expired bool `json:"expired"`

	// Used by Decode() method
	data []byte
}

func (model Target) New(data []byte) *Target {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Target) UnmarshalJSON(data []byte) error {
	type alias Target
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id           json.RawMessage `json:"$id"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		Name         json.RawMessage `json:"name"`
		UserId       json.RawMessage `json:"userId"`
		ProviderType json.RawMessage `json:"providerType"`
		Identifier   json.RawMessage `json:"identifier"`
		Expired      json.RawMessage `json:"expired"`
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
	if present.UserId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UserId); err != nil {
			return err
		}
	}
	if present.ProviderType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProviderType); err != nil {
			return err
		}
	}
	if present.Identifier == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Identifier); err != nil {
			return err
		}
	}
	if present.Expired == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Expired); err != nil {
			return err
		}
	}

	*model = Target(decoded)
	return nil
}

func (model *Target) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
