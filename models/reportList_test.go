package models

import (
    "encoding/json"
    "testing"
)

func TestReportListModel(t *testing.T) {
    model := ReportList{        Total: 5,        Reports: []Report{Report{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        AppId: "5e5ea5c16897e",        Type: "lighthouse",        Title: "Lighthouse audit for https://appwrite.io/",        Summary: "Performance score 78. 4 opportunities found.",        TargetType: "urls",        Target: "https://appwrite.io/",        Categories: []string{"test"},        Insights: []Insight{Insight{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        ReportId: "5e5ea5c16897e",        Type: "tablesDBIndex",        Severity: "warning",        Status: "active",        ResourceType: "databases",        ResourceId: "main",        ParentResourceType: "tables",        ParentResourceId: "orders",        Title: "Missing index on collection orders",        Summary: "Queries against `orders.status` are scanning the full collection.",        Ctas: []InsightCTA{InsightCTA{        Label: "Create missing index",        Service: "tablesDB",        Method: "createIndex",        Params: map[string]interface{}{},    },
            },    },
            },    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ReportList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
