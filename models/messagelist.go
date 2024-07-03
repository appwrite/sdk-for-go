package models


// MessageList Model
type MessageList struct {
    // Total number of messages documents that matched your query.
    Total int `json:"total"`
    // List of messages.
    Messages []interface{} `json:"messages"`

}
