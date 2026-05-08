package models

import (
    "encoding/json"
    "testing"
)

func TestPolicyPasswordPersonalDataModel(t *testing.T) {
    model := PolicyPasswordPersonalData{        Id: "password-dictionary",        Enabled: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PolicyPasswordPersonalData
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
