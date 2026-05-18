package models

import (
    "encoding/json"
    "testing"
)

func TestUsageGaugeModel(t *testing.T) {
    model := UsageGauge{        Metric: "users",        Value: 1500,        Time: "2026-04-09T12:00:00.000+00:00",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result UsageGauge
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Metric != model.Metric {
        t.Errorf("Expected Metric %v, got %v", model.Metric, result.Metric)
    }
    if result.Value != model.Value {
        t.Errorf("Expected Value %v, got %v", model.Value, result.Value)
    }
    if result.Time != model.Time {
        t.Errorf("Expected Time %v, got %v", model.Time, result.Time)
    }}
