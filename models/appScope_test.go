package models

import (
    "encoding/json"
    "testing"
)

func TestAppScopeModel(t *testing.T) {
    model := AppScope{        Value: "organization:organization.read",        Description: "Access to read the organization",        Type: "organization",        Category: "Organization",        Deprecated: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AppScope
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Value != model.Value {
        t.Errorf("Expected Value %v, got %v", model.Value, result.Value)
    }
    if result.Description != model.Description {
        t.Errorf("Expected Description %v, got %v", model.Description, result.Description)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.Category != model.Category {
        t.Errorf("Expected Category %v, got %v", model.Category, result.Category)
    }
    if result.Deprecated != model.Deprecated {
        t.Errorf("Expected Deprecated %v, got %v", model.Deprecated, result.Deprecated)
    }}
