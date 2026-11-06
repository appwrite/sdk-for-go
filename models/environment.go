package models

import (
	"encoding/json"
	"errors"
)

// Environment Model
type Environment struct {
	// Environment ID, as passed in the X-Appwrite-Environment header.
	Id string `json:"$id"`
	// Environment creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Environment update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// ID of the project the environment belongs to, as passed in the
	// X-Appwrite-Project header.
	ProjectId string `json:"projectId"`
	// Environment name.
	Name string `json:"name"`

	// Used by Decode() method
	data []byte
}

func (model Environment) New(data []byte) *Environment {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *Environment) UnmarshalJSON(data []byte) error {
	type alias Environment
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id        json.RawMessage `json:"$id"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UpdatedAt json.RawMessage `json:"$updatedAt"`
		ProjectId json.RawMessage `json:"projectId"`
		Name      json.RawMessage `json:"name"`
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
	if present.ProjectId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProjectId); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}

	*model = Environment(decoded)
	return nil
}

func (model *Environment) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
