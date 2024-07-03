package models


// TopicList Model
type TopicList struct {
    // Total number of topics documents that matched your query.
    Total int `json:"total"`
    // List of topics.
    Topics []interface{} `json:"topics"`

}
