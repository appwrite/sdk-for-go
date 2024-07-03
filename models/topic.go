package models


// Topic Model
type Topic struct {
    // Topic ID.
    Id string `json:"$id"`
    // Topic creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Topic update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // The name of the topic.
    Name string `json:"name"`
    // Total count of email subscribers subscribed to the topic.
    EmailTotal int `json:"emailTotal"`
    // Total count of SMS subscribers subscribed to the topic.
    SmsTotal int `json:"smsTotal"`
    // Total count of push subscribers subscribed to the topic.
    PushTotal int `json:"pushTotal"`
    // Subscribe permissions.
    Subscribe []interface{} `json:"subscribe"`

}
