package models


// AttributesList Model
type AttributeList struct {
    // Total number of attributes in the given collection.
    Total int `json:"total"`
    // List of attributes.
    Attributes []interface{} `json:"attributes"`

}
