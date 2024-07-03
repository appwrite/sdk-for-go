package models


// Message Model
type Message struct {
    // Message ID.
    Id string `json:"$id"`
    // Message creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Message update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Message provider type.
    ProviderType string `json:"providerType"`
    // Topic IDs set as recipients.
    Topics []interface{} `json:"topics"`
    // User IDs set as recipients.
    Users []interface{} `json:"users"`
    // Target IDs set as recipients.
    Targets []interface{} `json:"targets"`
    // The scheduled time for message.
    ScheduledAt string `json:"scheduledAt"`
    // The time when the message was delivered.
    DeliveredAt string `json:"deliveredAt"`
    // Delivery errors if any.
    DeliveryErrors []interface{} `json:"deliveryErrors"`
    // Number of recipients the message was delivered to.
    DeliveredTotal int `json:"deliveredTotal"`
    // Data of the message.
    Data interface{} `json:"data"`
    // Status of delivery.
    Status string `json:"status"`

}
