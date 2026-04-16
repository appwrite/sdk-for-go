package models

import (
    "encoding/json"
    "testing"
)

func TestHealthStatusListModel(t *testing.T) {
    model := HealthStatusList{        Total: 5,        Statuses: []HealthStatus{HealthStatus{        Name: "database",        Ping: 128,        Status: "pass",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result HealthStatusList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
