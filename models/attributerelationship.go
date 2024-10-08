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
    Type string `json:"xtype"`
    // Attribute status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Error message. Displays error generated on failure of creating or deleting
    // an attribute.
    Error string `json:"xerror"`
    // Is attribute required?
    Required bool `json:"required"`
    // Is attribute an array?
    Array bool `json:"array"`
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
