package models

import (
	"encoding/json"
	"errors"
)

// PlatformLinux Model
type PlatformLinux struct {
	// Platform ID.
	Id string `json:"$id"`
	// Platform creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Platform update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Platform name.
	Name string `json:"name"`
	// Platform type. Possible values are: windows, apple, android, linux, web.
	Type string `json:"type"`
	// Linux package name.
	PackageName string `json:"packageName"`

	// Used by Decode() method
	data []byte
}

func (model PlatformLinux) New(data []byte) *PlatformLinux {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *PlatformLinux) UnmarshalJSON(data []byte) error {
	type alias PlatformLinux
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id          json.RawMessage `json:"$id"`
		CreatedAt   json.RawMessage `json:"$createdAt"`
		UpdatedAt   json.RawMessage `json:"$updatedAt"`
		Name        json.RawMessage `json:"name"`
		PackageName json.RawMessage `json:"packageName"`
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
	if present.PackageName == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.PackageName); err != nil {
			return err
		}
	}

	*model = PlatformLinux(decoded)
	return nil
}

func (model *PlatformLinux) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
