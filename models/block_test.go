package models

import (
    "encoding/json"
    "testing"
)

func TestBlockModel(t *testing.T) {
    model := Block{        CreatedAt: "2020-10-15T06:38:00.000+00:00",        ResourceType: "project",        ResourceId: "5e5ea5c16897e",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Block
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.ResourceType != model.ResourceType {
        t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
    }
    if result.ResourceId != model.ResourceId {
        t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
    }}
