package models

import (
    "encoding/json"
    "errors"
)

// Replicas Model
type DedicatedDatabaseReplicas struct {
    // Number of configured replicas. Zero means high availability is disabled.
    Replicas int `json:"replicas"`
    // Replication sync mode. Possible values: async (asynchronous, fastest), sync
    // (synchronous, strong consistency), quorum (quorum-based, majority of
    // replicas must confirm).
    SyncMode string `json:"syncMode"`
    // Per-pod statuses for the primary and every replica.
    Members []DedicatedDatabaseMember `json:"members"`

    // Used by Decode() method
    data []byte
}

func (model DedicatedDatabaseReplicas) New(data []byte) *DedicatedDatabaseReplicas {
    model.data = data
    return &model
}

func (model *DedicatedDatabaseReplicas) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}