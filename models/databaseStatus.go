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
    // Database engine: postgresql, mysql, mariadb, or mongodb.
    Engine string `json:"engine"`
    // Database engine version.
    Version string `json:"version"`
    // Database uptime in seconds.
    Uptime int `json:"uptime"`
    // Connection statistics.
    Connections DatabaseStatusConnections `json:"connections"`
    // List of database replicas and their status.
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