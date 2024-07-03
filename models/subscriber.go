package models


// Subscriber Model
type Subscriber struct {
    // Subscriber ID.
    Id string `json:"$id"`
    // Subscriber creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Subscriber update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Target ID.
    TargetId string `json:"targetId"`
    // Target.
    Target interface{} `json:"target"`
    // Topic ID.
    UserId string `json:"userId"`
    // User Name.
    UserName string `json:"userName"`
    // Topic ID.
    TopicId string `json:"topicId"`
    // The target provider type. Can be one of the following: `email`, `sms` or
    // `push`.
    ProviderType string `json:"providerType"`

}
