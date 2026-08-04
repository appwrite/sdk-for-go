package models

import (
    "encoding/json"
    "errors"
)

// Status Model
type DatabaseStatus struct {
    // Overall health status: healthy, degraded, or unhealthy.
    Health string `json:"health"`
    // Whether the database is ready to accept connections.
    Ready bool `json:"ready"`
    // Database engine: postgresql, mysql, or mongodb.
    Engine string `json:"engine"`
    // Database engine version.
    Version string `json:"version"`
    // Database uptime in seconds.
    Uptime int `json:"uptime"`
    // Connection statistics.
    Connections DatabaseStatusConnections `json:"connections"`
    // Requested replication sync mode. Possible values: async, sync, quorum.
    // Compare with effectiveSyncMode for what the primary is enforcing.
    SyncMode string `json:"syncMode"`
    // Replication sync mode the primary is actually enforcing. Null when high
    // availability is disabled or the state could not be read.
    EffectiveSyncMode string `json:"effectiveSyncMode"`
    // Whether the enforced replication is weaker than the requested syncMode.
    SyncDegraded bool `json:"syncDegraded"`
    // Number of standby acknowledgements the primary waits for before a write is
    // committed.
    SyncAcknowledgements int `json:"syncAcknowledgements"`
    // Number of standbys registered with the primary for synchronous replication.
    SyncStandbyCount int `json:"syncStandbyCount"`
    // Whether the reported sync state was read from the engine. When false the
    // state could not be confirmed and the other sync fields carry no reading.
    SyncStateConfirmed bool `json:"syncStateConfirmed"`
    // List of database replicas and their status. Every configured member
    // appears, including one the backend has not brought up, which is reported as
    // not healthy.
    Replicas []DatabaseStatusReplica `json:"replicas"`
    // Storage volume information.
    Volumes []DatabaseStatusVolume `json:"volumes"`

    // Used by Decode() method
    data []byte
}

func (model DatabaseStatus) New(data []byte) *DatabaseStatus {
    model.data = data
    return &model
}

func (model *DatabaseStatus) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}