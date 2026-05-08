package models

import (
    "encoding/json"
    "testing"
)

func TestPolicyUserLimitModel(t *testing.T) {
    model := PolicyUserLimit{        Id: "password-dictionary",        Total: 100,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PolicyUserLimit
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
