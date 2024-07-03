package models


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

}
