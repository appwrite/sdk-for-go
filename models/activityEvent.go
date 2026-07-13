package models

import (
    "encoding/json"
    "errors"
)

// ActivityEvent Model
type ActivityEvent struct {
    // Event ID.
    Id string `json:"$id"`
    // Actor type.
    ActorType string `json:"actorType"`
    // Actor ID.
    ActorId string `json:"actorId"`
    // Actor Email.
    ActorEmail string `json:"actorEmail"`
    // Actor Name.
    ActorName string `json:"actorName"`
    // Resource parent.
    ResourceParent string `json:"resourceParent"`
    // Resource type.
    ResourceType string `json:"resourceType"`
    // Resource ID.
    ResourceId string `json:"resourceId"`
    // Resource.
    Resource string `json:"resource"`
    // Event name.
    Event string `json:"event"`
    // User agent.
    UserAgent string `json:"userAgent"`
    // IP address.
    Ip string `json:"ip"`
    // API mode when event triggered.
    Mode string `json:"mode"`
    // Location.
    Country string `json:"country"`
    // Log creation date in ISO 8601 format.
    Time string `json:"time"`
    // Project ID.
    ProjectId string `json:"projectId"`
    // Team ID.
    TeamId string `json:"teamId"`
    // Hostname.
    Hostname string `json:"hostname"`

    // Used by Decode() method
    data []byte
}

func (model ActivityEvent) New(data []byte) *ActivityEvent {
    model.data = data
    return &model
}

func (model *ActivityEvent) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}