package models

import (
    "encoding/json"
    "testing"
)

func TestDedicatedDatabaseReplicasModel(t *testing.T) {
    model := DedicatedDatabaseReplicas{        Replicas: 2,        SyncMode: "async",        Members: []DedicatedDatabaseMember{DedicatedDatabaseMember{        Id: "1",        Role: "replica",        Status: "active",        LagSeconds: 0.5,    },
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
    }}
