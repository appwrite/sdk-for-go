package models

import (
    "encoding/json"
    "testing"
)

func TestUsageEventListModel(t *testing.T) {
    model := UsageEventList{        Total: 5,        Events: []UsageEvent{UsageEvent{        Metric: "bandwidth",        Value: 5000,        Time: "2026-04-09T12:00:00.000+00:00",        Path: "/v1/storage/files",        Method: "POST",        Status: "201",        ResourceType: "bucket",        ResourceId: "abc123",        CountryCode: "US",        UserAgent: "AppwriteSDK/1.0",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result UsageEventList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
