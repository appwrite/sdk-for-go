package models


// SessionsList Model
type SessionList struct {
    // Total number of sessions documents that matched your query.
    Total int `json:"total"`
    // List of sessions.
    Sessions []interface{} `json:"sessions"`

}
