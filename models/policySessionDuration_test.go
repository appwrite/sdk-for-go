package models

import (
    "encoding/json"
    "testing"
)

func TestPolicySessionDurationModel(t *testing.T) {
    model := PolicySessionDuration{        Id: "password-dictionary",        Duration: 3600,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PolicySessionDuration
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Duration != model.Duration {
        t.Errorf("Expected Duration %v, got %v", model.Duration, result.Duration)
    }}
