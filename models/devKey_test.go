package models

import (
    "encoding/json"
    "testing"
)

func TestDevKeyModel(t *testing.T) {
    model := DevKey{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Dev API Key",        Expire: "2020-10-15T06:38:00.000+00:00",        Secret: "919c2d18fb5d4...a2ae413da83346ad2",        AccessedAt: "2020-10-15T06:38:00.000+00:00",        Sdks: []string{"test"},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DevKey
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Expire != model.Expire {
        t.Errorf("Expected Expire %v, got %v", model.Expire, result.Expire)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }
    if result.AccessedAt != model.AccessedAt {
        t.Errorf("Expected AccessedAt %v, got %v", model.AccessedAt, result.AccessedAt)
    }}
