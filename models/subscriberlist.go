package models


// SubscriberList Model
type SubscriberList struct {
    // Total number of subscribers documents that matched your query.
    Total int `json:"total"`
    // List of subscribers.
    Subscribers []interface{} `json:"subscribers"`

}
