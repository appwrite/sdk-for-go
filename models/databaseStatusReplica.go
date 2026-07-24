package models

import (
    "encoding/json"
    "errors"
)

// Replica Model
type DatabaseStatusReplica struct {
    // StatefulSet pod index (0 = primary, 1+ = replicas).
    Index int `json:"index"`
    // Replica role: primary or replica.
    Role string `json:"role"`
    // Whether the replica is healthy.
    Healthy bool `json:"healthy"`
    // Replication lag in seconds (null for primary).
    LagSeconds float64 `json:"lagSeconds"`

    // Used by Decode() method
    data []byte
}

func (model DatabaseStatusReplica) New(data []byte) *DatabaseStatusReplica {
    model.data = data
    return &model
}

func (model *DatabaseStatusReplica) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}