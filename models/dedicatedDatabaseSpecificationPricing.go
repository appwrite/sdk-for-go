package models

import (
    "encoding/json"
    "errors"
)

// SpecificationPricing Model
type DedicatedDatabaseSpecificationPricing struct {
    // Price per GB of storage above the included amount, per month, in USD.
    StorageOverageRate float64 `json:"storageOverageRate"`
    // Price per GB of bandwidth above the included amount, per month, in USD.
    BandwidthOverageRate float64 `json:"bandwidthOverageRate"`
    // High availability replica price as a fraction of the specification cost.
    ReplicaRate float64 `json:"replicaRate"`
    // Cross-region replica price as a fraction of the specification cost.
    CrossRegionReplicaRate float64 `json:"crossRegionReplicaRate"`
    // Point-in-time recovery price as a fraction of the specification cost.
    PitrRate float64 `json:"pitrRate"`

    // Used by Decode() method
    data []byte
}

func (model DedicatedDatabaseSpecificationPricing) New(data []byte) *DedicatedDatabaseSpecificationPricing {
    model.data = data
    return &model
}

func (model *DedicatedDatabaseSpecificationPricing) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}