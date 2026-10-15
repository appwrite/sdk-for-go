package models

import (
	"encoding/json"
	"errors"
)

// AttributeRelationship Model
type AttributeRelationship struct {
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
	// The ID of the related collection.
	RelatedCollection string `json:"relatedCollection"`
	// The type of the relationship.
	RelationType string `json:"relationType"`
	// Is the relationship two-way?
	TwoWay bool `json:"twoWay"`
	// The key of the two-way relationship.
	TwoWayKey string `json:"twoWayKey"`
	// How deleting the parent document will propagate to child documents.
	OnDelete string `json:"onDelete"`
	// Whether this is the parent or child side of the relationship
	Side string `json:"side"`

	// Used by Decode() method
	data []byte
}

func (model AttributeRelationship) New(data []byte) *AttributeRelationship {
	model.data = data
	return &model
}

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *AttributeRelationship) UnmarshalJSON(data []byte) error {
	type alias AttributeRelationship
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Key          json.RawMessage `json:"key"`
		Type         json.RawMessage `json:"type"`
		Error        json.RawMessage `json:"error"`
		Required     json.RawMessage `json:"required"`
		CreatedAt    json.RawMessage `json:"$createdAt"`
		UpdatedAt    json.RawMessage `json:"$updatedAt"`
		RelationType json.RawMessage `json:"relationType"`
		TwoWay       json.RawMessage `json:"twoWay"`
		TwoWayKey    json.RawMessage `json:"twoWayKey"`
		OnDelete     json.RawMessage `json:"onDelete"`
		Side         json.RawMessage `json:"side"`
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
	if present.RelationType == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.RelationType); err != nil {
			return err
		}
	}
	if present.TwoWay == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.TwoWay); err != nil {
			return err
		}
	}
	if present.TwoWayKey == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.TwoWayKey); err != nil {
			return err
		}
	}
	if present.OnDelete == nil {
		if err := json.Unmarshal([]byte("\"restrict\""), &decoded.OnDelete); err != nil {
			return err
		}
	}
	if present.Side == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Side); err != nil {
			return err
		}
	}

	*model = AttributeRelationship(decoded)
	return nil
}

func (model *AttributeRelationship) Decode(value interface{}) error {
	if len(model.data) <= 0 {
		return errors.New("method Decode() cannot be used on nested struct")
	}

	err := json.Unmarshal(model.data, value)
	if err != nil {
		return err
	}

	return nil
}
