package models

import (
    "encoding/json"
    "testing"
)

func TestDedicatedDatabaseSpecificationPricingModel(t *testing.T) {
    model := DedicatedDatabaseSpecificationPricing{        StorageOverageRate: 0.125,        BandwidthOverageRate: 0.08,        ReplicaRate: 1,        CrossRegionReplicaRate: 1,        PitrRate: 0.2,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DedicatedDatabaseSpecificationPricing
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.StorageOverageRate != model.StorageOverageRate {
        t.Errorf("Expected StorageOverageRate %v, got %v", model.StorageOverageRate, result.StorageOverageRate)
    }
    if result.BandwidthOverageRate != model.BandwidthOverageRate {
        t.Errorf("Expected BandwidthOverageRate %v, got %v", model.BandwidthOverageRate, result.BandwidthOverageRate)
    }
    if result.ReplicaRate != model.ReplicaRate {
        t.Errorf("Expected ReplicaRate %v, got %v", model.ReplicaRate, result.ReplicaRate)
    }
    if result.CrossRegionReplicaRate != model.CrossRegionReplicaRate {
        t.Errorf("Expected CrossRegionReplicaRate %v, got %v", model.CrossRegionReplicaRate, result.CrossRegionReplicaRate)
    }
    if result.PitrRate != model.PitrRate {
        t.Errorf("Expected PitrRate %v, got %v", model.PitrRate, result.PitrRate)
    }}
