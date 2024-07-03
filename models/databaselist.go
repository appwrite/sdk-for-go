package models


// DatabasesList Model
type DatabaseList struct {
    // Total number of databases documents that matched your query.
    Total int `json:"total"`
    // List of databases.
    Databases []interface{} `json:"databases"`

}
