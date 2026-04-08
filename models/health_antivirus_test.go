package models

import (
    "encoding/json"
    "testing"
)

func TestHealthAntivirusModel(t *testing.T) {
    model := HealthAntivirus{        Version: "1.0.0",        Status: "online",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result HealthAntivirus
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Version != model.Version {
        t.Errorf("Expected Version %v, got %v", model.Version, result.Version)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }}
