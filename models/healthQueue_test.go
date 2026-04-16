package models

import (
    "encoding/json"
    "testing"
)

func TestHealthQueueModel(t *testing.T) {
    model := HealthQueue{        Size: 8,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result HealthQueue
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Size != model.Size {
        t.Errorf("Expected Size %v, got %v", model.Size, result.Size)
    }}
