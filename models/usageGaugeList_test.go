package models

import (
    "encoding/json"
    "testing"
)

func TestUsageGaugeListModel(t *testing.T) {
    model := UsageGaugeList{        Total: 5,        Gauges: []UsageGauge{UsageGauge{        Metric: "users",        Value: 1500,        Time: "2026-04-09T12:00:00.000+00:00",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result UsageGaugeList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
