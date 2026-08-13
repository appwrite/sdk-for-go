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
    // Whether the other sync fields are an engine reading rather than a recorded
    // estimate. True when the primary answered what it is enforcing, including
    // when that answer contradicted the record, in which case the contradicted
    // values are replaced by the ones the engine reports. False when the reading
    // could not be taken: the probe did not answer, there was no engine to ask,
    // or the values describe a configuration change just applied rather than
    // anything measured. Absent when no engine was asked at all, so an unprobed
    // database is distinguishable from an unconfirmed one. False never means a
    // standby was found lagging, because it is the absence of a reading rather
    // than a negative one, so draw no conclusion about replication health from it
    // or from a response that omits it.
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