package models

import (
    "encoding/json"
    "testing"
)

func TestInsightModel(t *testing.T) {
    model := Insight{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        ReportId: "5e5ea5c16897e",        Type: "tablesDBIndex",        Severity: "warning",        Status: "active",        ResourceType: "databases",        ResourceId: "main",        ParentResourceType: "tables",        ParentResourceId: "orders",        Title: "Missing index on collection orders",        Summary: "Queries against `orders.status` are scanning the full collection.",        Ctas: []InsightCTA{InsightCTA{        Label: "Create missing index",        Service: "tablesDB",        Method: "createIndex",        Params: map[string]interface{}{},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Insight
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
    if result.ReportId != model.ReportId {
        t.Errorf("Expected ReportId %v, got %v", model.ReportId, result.ReportId)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.Severity != model.Severity {
        t.Errorf("Expected Severity %v, got %v", model.Severity, result.Severity)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.ResourceType != model.ResourceType {
        t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
    }
    if result.ResourceId != model.ResourceId {
        t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
    }
    if result.ParentResourceType != model.ParentResourceType {
        t.Errorf("Expected ParentResourceType %v, got %v", model.ParentResourceType, result.ParentResourceType)
    }
    if result.ParentResourceId != model.ParentResourceId {
        t.Errorf("Expected ParentResourceId %v, got %v", model.ParentResourceId, result.ParentResourceId)
    }
    if result.Title != model.Title {
        t.Errorf("Expected Title %v, got %v", model.Title, result.Title)
    }
    if result.Summary != model.Summary {
        t.Errorf("Expected Summary %v, got %v", model.Summary, result.Summary)
    }}
