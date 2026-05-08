package models

import (
    "encoding/json"
    "testing"
)

func TestPolicySessionInvalidationModel(t *testing.T) {
    model := PolicySessionInvalidation{        Id: "password-dictionary",        Enabled: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PolicySessionInvalidation
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }}
