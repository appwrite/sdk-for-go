package models

import (
    "encoding/json"
    "errors"
)

// ColumnRelationship Model
type ColumnRelationship struct {
    // Column Key.
    Key string `json:"key"`
    // Column type.
    Type string `json:"type"`
    // Column status. Possible values: `available`, `processing`, `deleting`,
    // `stuck`, or `failed`
    Status string `json:"status"`
    // Error message. Displays error generated on failure of creating or deleting
    // an column.
    Error string `json:"error"`
    // Is column required?
    Required bool `json:"required"`
    // Is column an array?
    Array bool `json:"array"`
    // Column creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Column update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // The ID of the related table.
    RelatedTable string `json:"relatedTable"`
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

func (model ColumnRelationship) New(data []byte) *ColumnRelationship {
    model.data = data
    return &model
}

func (model *ColumnRelationship) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}