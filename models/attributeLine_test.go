package models

import (
    "encoding/json"
    "testing"
)

func TestAttributeLineModel(t *testing.T) {
    model := AttributeLine{        Key: "fullName",        Type: "string",        Status: "available",        Error: "string",        Required: true,        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AttributeLine
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Key != model.Key {
        t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.Error != model.Error {
        t.Errorf("Expected Error %v, got %v", model.Error, result.Error)
    }
    if result.Required != model.Required {
        t.Errorf("Expected Required %v, got %v", model.Required, result.Required)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }}
