package models

import (
    "encoding/json"
    "testing"
)

func TestMessageModel(t *testing.T) {
    model := Message{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        ProviderType: "email",        Topics: []string{"test"},        Users: []string{"test"},        Targets: []string{"test"},        DeliveredTotal: 1,        Data: map[string]interface{}{},        Status: "processing",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Message
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
    if result.ProviderType != model.ProviderType {
        t.Errorf("Expected ProviderType %v, got %v", model.ProviderType, result.ProviderType)
    }
    if result.DeliveredTotal != model.DeliveredTotal {
        t.Errorf("Expected DeliveredTotal %v, got %v", model.DeliveredTotal, result.DeliveredTotal)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }}
