package models

import (
    "encoding/json"
    "testing"
)

func TestDedicatedDatabaseReplicasModel(t *testing.T) {
    model := DedicatedDatabaseReplicas{        Replicas: 2,        SyncMode: "async",        SyncDegraded: true,        SyncAcknowledgements: 1,        SyncStandbyCount: 2,        SyncStateConfirmed: true,        Members: []DedicatedDatabaseMember{DedicatedDatabaseMember{        Id: "1",        Role: "replica",        Status: "active",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DedicatedDatabaseReplicas
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Replicas != model.Replicas {
        t.Errorf("Expected Replicas %v, got %v", model.Replicas, result.Replicas)
    }
    if result.SyncMode != model.SyncMode {
        t.Errorf("Expected SyncMode %v, got %v", model.SyncMode, result.SyncMode)
    }
    if result.SyncDegraded != model.SyncDegraded {
        t.Errorf("Expected SyncDegraded %v, got %v", model.SyncDegraded, result.SyncDegraded)
    }
    if result.SyncAcknowledgements != model.SyncAcknowledgements {
        t.Errorf("Expected SyncAcknowledgements %v, got %v", model.SyncAcknowledgements, result.SyncAcknowledgements)
    }
    if result.SyncStandbyCount != model.SyncStandbyCount {
        t.Errorf("Expected SyncStandbyCount %v, got %v", model.SyncStandbyCount, result.SyncStandbyCount)
    }
    if result.SyncStateConfirmed != model.SyncStateConfirmed {
        t.Errorf("Expected SyncStateConfirmed %v, got %v", model.SyncStateConfirmed, result.SyncStateConfirmed)
    }}
