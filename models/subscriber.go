package models

import (
    "encoding/json"
    "errors"
)

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
    Target Target `json:"target"`
    // Topic ID.
    UserId string `json:"userId"`
    // User Name.
    UserName string `json:"userName"`
    // Topic ID.
    TopicId string `json:"topicId"`
    // The target provider type. Can be one of the following: `email`, `sms` or
    // `push`.
    ProviderType string `json:"providerType"`

    // Used by Decode() method
    data []byte
}

func (model Subscriber) New(data []byte) *Subscriber {
    model.data = data
    return &model
}

func (model *Subscriber) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
