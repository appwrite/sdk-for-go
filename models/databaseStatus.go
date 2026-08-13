package models

import (
    "encoding/json"
    "errors"
)

// Status Model
type DatabaseStatus struct {
    // Overall health status: healthy, degraded, unhealthy, or unknown when
    // nothing could be measured.
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