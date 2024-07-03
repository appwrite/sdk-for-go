package models


// DocumentsList Model
type DocumentList struct {
    // Total number of documents documents that matched your query.
    Total int `json:"total"`
    // List of documents.
    Documents []interface{} `json:"documents"`

}
