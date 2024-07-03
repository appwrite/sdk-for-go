package models


// BucketsList Model
type BucketList struct {
    // Total number of buckets documents that matched your query.
    Total int `json:"total"`
    // List of buckets.
    Buckets []interface{} `json:"buckets"`

}
