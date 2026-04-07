package models

import (
    "encoding/json"
    "testing"
)

func TestDocumentModel(t *testing.T) {
    model := Document{        Id: "5e5ea5c16897e",        Sequence: "1",        CollectionId: "5e5ea5c15117e",        DatabaseId: "5e5ea5c15117e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Document
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Sequence != model.Sequence {
        t.Errorf("Expected Sequence %v, got %v", model.Sequence, result.Sequence)
    }
    if result.CollectionId != model.CollectionId {
        t.Errorf("Expected CollectionId %v, got %v", model.CollectionId, result.CollectionId)
    }
    if result.DatabaseId != model.DatabaseId {
        t.Errorf("Expected DatabaseId %v, got %v", model.DatabaseId, result.DatabaseId)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }}
