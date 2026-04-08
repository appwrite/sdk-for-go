package models

import (
    "encoding/json"
    "testing"
)

func TestTransactionModel(t *testing.T) {
    model := Transaction{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Status: "pending",        Operations: 5,        ExpiresAt: "2020-10-15T06:38:00.000+00:00",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Transaction
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
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.Operations != model.Operations {
        t.Errorf("Expected Operations %v, got %v", model.Operations, result.Operations)
    }
    if result.ExpiresAt != model.ExpiresAt {
        t.Errorf("Expected ExpiresAt %v, got %v", model.ExpiresAt, result.ExpiresAt)
    }}
