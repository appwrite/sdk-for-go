package models

import (
    "encoding/json"
    "testing"
)

func TestDedicatedDatabaseOperationModel(t *testing.T) {
    model := DedicatedDatabaseOperation{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        DatabaseId: "5e5ea5c16897e",        Type: "update",        Status: "completed",        Attempts: 1,        ErrorCode: "Interrupted",        ErrorMessage: "string",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DedicatedDatabaseOperation
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.DatabaseId != model.DatabaseId {
        t.Errorf("Expected DatabaseId %v, got %v", model.DatabaseId, result.DatabaseId)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.Attempts != model.Attempts {
        t.Errorf("Expected Attempts %v, got %v", model.Attempts, result.Attempts)
    }
    if result.ErrorCode != model.ErrorCode {
        t.Errorf("Expected ErrorCode %v, got %v", model.ErrorCode, result.ErrorCode)
    }
    if result.ErrorMessage != model.ErrorMessage {
        t.Errorf("Expected ErrorMessage %v, got %v", model.ErrorMessage, result.ErrorMessage)
    }}
