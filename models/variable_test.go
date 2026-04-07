package models

import (
    "encoding/json"
    "testing"
)

func TestVariableModel(t *testing.T) {
    model := Variable{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "API_KEY",        Value: "myPa$$word1",        Secret: true,        ResourceType: "function",        ResourceId: "myAwesomeFunction",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Variable
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
    if result.Key != model.Key {
        t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
    }
    if result.Value != model.Value {
        t.Errorf("Expected Value %v, got %v", model.Value, result.Value)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }
    if result.ResourceType != model.ResourceType {
        t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
    }
    if result.ResourceId != model.ResourceId {
        t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
    }}
