package models

import (
    "encoding/json"
    "errors"
)

// Member Model
type DedicatedDatabaseMember struct {
    // Member identifier.
    Id string `json:"$id"`
    // Member role. Possible values: primary (accepts reads and writes), replica
    // (read-only follower).
    Role string `json:"role"`
    // Member pod status. Possible values: provisioning (pod missing or Pending),
    // starting (Running but not Ready), active (Running and Ready), failed
    // (Failed phase or CrashLoopBackOff container), or the lowercased pod phase
    // reported by the cluster.
    Status string `json:"status"`
    // Replication lag in seconds.
    LagSeconds float64 `json:"lagSeconds"`

    // Used by Decode() method
    data []byte
}

func (model DedicatedDatabaseMember) New(data []byte) *DedicatedDatabaseMember {
    model.data = data
    return &model
}

func (model *DedicatedDatabaseMember) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}