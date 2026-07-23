package models

import (
    "encoding/json"
    "testing"
)

func TestDedicatedDatabaseSpecificationListModel(t *testing.T) {
    model := DedicatedDatabaseSpecificationList{        Specifications: []DedicatedDatabaseSpecification{DedicatedDatabaseSpecification{        Slug: "s-2vcpu-2gb",        Name: "Standard",        Price: 20,        Cpu: 2000,        Memory: 2048,        MaxConnections: 200,        IncludedStorage: 25,        IncludedBandwidth: 200,        Enabled: true,    },
            },        Total: 9,        Pricing: DedicatedDatabaseSpecificationPricing{        StorageOverageRate: 0.125,        BandwidthOverageRate: 0.08,        ReplicaRate: 1,        CrossRegionReplicaRate: 1,        PitrRate: 0.2,    },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DedicatedDatabaseSpecificationList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
