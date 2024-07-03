package models


// CollectionsList Model
type CollectionList struct {
    // Total number of collections documents that matched your query.
    Total int `json:"total"`
    // List of collections.
    Collections []interface{} `json:"collections"`

}
