package models

import (
    "encoding/json"
    "testing"
)

func TestPolicyPasswordHistoryModel(t *testing.T) {
    model := PolicyPasswordHistory{        Id: "password-dictionary",        Total: 5,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PolicyPasswordHistory
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
