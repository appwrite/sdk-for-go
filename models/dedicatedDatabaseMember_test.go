package models

import (
    "encoding/json"
    "testing"
)

func TestDedicatedDatabaseMemberModel(t *testing.T) {
    model := DedicatedDatabaseMember{        Id: "1",        Role: "replica",        Status: "active",        LagSeconds: 0.5,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DedicatedDatabaseMember
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Role != model.Role {
        t.Errorf("Expected Role %v, got %v", model.Role, result.Role)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.LagSeconds != model.LagSeconds {
        t.Errorf("Expected LagSeconds %v, got %v", model.LagSeconds, result.LagSeconds)
    }}
