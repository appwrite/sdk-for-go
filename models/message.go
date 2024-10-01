package models

import (
    "encoding/json"
    "errors"
)

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
    Topics []string `json:"topics"`
    // User IDs set as recipients.
    Users []string `json:"users"`
    // Target IDs set as recipients.
    Targets []string `json:"targets"`
    // The scheduled time for message.
    ScheduledAt string `json:"scheduledAt"`
    // The time when the message was delivered.
    DeliveredAt string `json:"deliveredAt"`
    // Delivery errors if any.
    DeliveryErrors []string `json:"deliveryErrors"`
    // Number of recipients the message was delivered to.
    DeliveredTotal int `json:"deliveredTotal"`
    // Data of the message.
    Data interface{} `json:"data"`
    // Status of delivery.
    Status string `json:"status"`

    // Used by Decode() method
    data []byte
}

func (model Message) New(data []byte) *Message {
    model.data = data
    return &model
}

func (model *Message) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
