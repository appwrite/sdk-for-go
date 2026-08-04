package models

import (
    "encoding/json"
    "errors"
)

// Replicas Model
type DedicatedDatabaseReplicas struct {
    // Number of configured replicas. Zero means high availability is disabled.
    Replicas int `json:"replicas"`
    // Requested replication sync mode. Possible values: async (asynchronous,
    // fastest), sync (synchronous, strong consistency), quorum (quorum-based,
    // majority of replicas must confirm). This is what was asked for; compare it
    // with effectiveSyncMode for what the primary is enforcing.
    SyncMode string `json:"syncMode"`
    // Replication sync mode the primary is actually enforcing. Null when high
    // availability is disabled or the state could not be read. A value below the
    // requested syncMode means writes are being acknowledged with weaker
    // durability than configured.
    EffectiveSyncMode string `json:"effectiveSyncMode"`
    // Whether the enforced replication is weaker than the requested syncMode.
    SyncDegraded bool `json:"syncDegraded"`
    // Number of standby acknowledgements the primary waits for before a write is
    // committed. Zero means writes are acknowledged locally.
    SyncAcknowledgements int `json:"syncAcknowledgements"`
    // Number of standbys registered with the primary for synchronous replication.
    SyncStandbyCount int `json:"syncStandbyCount"`
    // Whether the reported sync state was read from the engine. When false the
    // state could not be confirmed and the other sync fields carry no reading.
    SyncStateConfirmed bool `json:"syncStateConfirmed"`
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