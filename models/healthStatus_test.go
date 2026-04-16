package models

import (
    "encoding/json"
    "testing"
)

func TestHealthStatusModel(t *testing.T) {
    model := HealthStatus{        Name: "database",        Ping: 128,        Status: "pass",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result HealthStatus
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Ping != model.Ping {
        t.Errorf("Expected Ping %v, got %v", model.Ping, result.Ping)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }}
