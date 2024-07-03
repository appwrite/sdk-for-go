package models


// AttributeDatetime Model
type AttributeDatetime struct {
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
    // ISO 8601 format.
    Format string `json:"format"`
    // Default value for attribute when not provided. Only null is optional
    Default string `json:"xdefault"`

}
