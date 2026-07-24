package models

import (
    "encoding/json"
    "testing"
)

func TestDatabaseStatusModel(t *testing.T) {
    model := DatabaseStatus{        Health: "healthy",        Ready: true,        Engine: "postgresql",        Version: "17",        Uptime: 86400,        Connections: DatabaseStatusConnections{        Current: 12,        Max: 100,    },        Replicas: []DatabaseStatusReplica{DatabaseStatusReplica{        Index: 0,        Role: "primary",        Healthy: true,    },
            },        Volumes: []DatabaseStatusVolume{DatabaseStatusVolume{        Path: "/var/lib/postgresql/data",        UsedPercent: "45%",        Available: "55GB",        Mounted: true,    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DatabaseStatus
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Health != model.Health {
        t.Errorf("Expected Health %v, got %v", model.Health, result.Health)
    }
    if result.Ready != model.Ready {
        t.Errorf("Expected Ready %v, got %v", model.Ready, result.Ready)
    }
    if result.Engine != model.Engine {
        t.Errorf("Expected Engine %v, got %v", model.Engine, result.Engine)
    }
    if result.Version != model.Version {
        t.Errorf("Expected Version %v, got %v", model.Version, result.Version)
    }
    if result.Uptime != model.Uptime {
        t.Errorf("Expected Uptime %v, got %v", model.Uptime, result.Uptime)
    }}
