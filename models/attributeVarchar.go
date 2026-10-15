package models

import (
	"encoding/json"
	"errors"
)

// AttributeVarchar Model
type AttributeVarchar struct {
	// Attribute Key.
	Key string `json:"key"`
	// Attribute type.
	Type string `json:"type"`
	// Attribute status. Possible values: `available`, `processing`, `deleting`,
	// `stuck`, or `failed`
	Status string `json:"status"`
	// Error message. Displays error generated on failure of creating or deleting
	// an attribute.
	Error string `json:"error"`
	// Is attribute required?
	Required bool `json:"required"`
	// Is attribute an array?
	Array *bool `json:"array"`
	// Attribute creation date in ISO 8601 format.
	CreatedAt string `json:"$createdAt"`
	// Attribute update date in ISO 8601 format.
	UpdatedAt string `json:"$updatedAt"`
	// Attribute size.
	Size int `json:"size"`
	// Default value for attribute when not provided. Cannot be set when attribute
	// is required.
	Default *string `json:"default"`
	// Defines whether this attribute is encrypted or not.
	Encrypt *bool `json:"encrypt"`

	// Used by Decode() method
	data []byte
}

func (model AttributeVarchar) New(data []byte) *AttributeVarchar {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AttributeVarchar) UnmarshalJSON(data []byte) error {
	type alias AttributeVarchar
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Key       json.RawMessage `json:"key"`
		Type      json.RawMessage `json:"type"`
		Error     json.RawMessage `json:"error"`
		Required  json.RawMessage `json:"required"`
		CreatedAt json.RawMessage `json:"$createdAt"`
		UpdatedAt json.RawMessage `json:"$updatedAt"`
		Size      json.RawMessage `json:"size"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Key == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Key); err != nil {
			return err
		}
	}
	if present.Type == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Type); err != nil {
			return err
		}
	}
	if present.Error == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Error); err != nil {
			return err
		}
	}
	if present.Required == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Required); err != nil {
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
	if present.Size == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Size); err != nil {
			return err
		}
	}

	*model = AttributeVarchar(decoded)
	return nil
}

func (model *AttributeVarchar) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
