package models

import (
    "encoding/json"
    "testing"
)

func TestHealthTimeModel(t *testing.T) {
    model := HealthTime{        RemoteTime: 1639490751,        LocalTime: 1639490844,        Diff: 93,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result HealthTime
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.RemoteTime != model.RemoteTime {
        t.Errorf("Expected RemoteTime %v, got %v", model.RemoteTime, result.RemoteTime)
    }
    if result.LocalTime != model.LocalTime {
        t.Errorf("Expected LocalTime %v, got %v", model.LocalTime, result.LocalTime)
    }
    if result.Diff != model.Diff {
        t.Errorf("Expected Diff %v, got %v", model.Diff, result.Diff)
    }}
