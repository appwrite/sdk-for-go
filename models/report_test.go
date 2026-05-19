package models

import (
    "encoding/json"
    "testing"
)

func TestReportModel(t *testing.T) {
    model := Report{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        AppId: "5e5ea5c16897e",        Type: "lighthouse",        Title: "Lighthouse audit for https://appwrite.io/",        Summary: "Performance score 78. 4 opportunities found.",        TargetType: "urls",        Target: "https://appwrite.io/",        Categories: []string{"test"},        Insights: []Insight{Insight{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        ReportId: "5e5ea5c16897e",        Type: "tablesDBIndex",        Severity: "warning",        Status: "active",        ResourceType: "databases",        ResourceId: "main",        ParentResourceType: "tables",        ParentResourceId: "orders",        Title: "Missing index on collection orders",        Summary: "Queries against `orders.status` are scanning the full collection.",        Ctas: []InsightCTA{InsightCTA{        Label: "Create missing index",        Service: "tablesDB",        Method: "createIndex",        Params: map[string]interface{}{},    },
            },    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Report
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
    if result.AppId != model.AppId {
        t.Errorf("Expected AppId %v, got %v", model.AppId, result.AppId)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.Title != model.Title {
        t.Errorf("Expected Title %v, got %v", model.Title, result.Title)
    }
    if result.Summary != model.Summary {
        t.Errorf("Expected Summary %v, got %v", model.Summary, result.Summary)
    }
    if result.TargetType != model.TargetType {
        t.Errorf("Expected TargetType %v, got %v", model.TargetType, result.TargetType)
    }
    if result.Target != model.Target {
        t.Errorf("Expected Target %v, got %v", model.Target, result.Target)
    }}
