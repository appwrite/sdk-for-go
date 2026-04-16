package models

import (
    "encoding/json"
    "testing"
)

func TestAuthProviderModel(t *testing.T) {
    model := AuthProvider{        Key: "github",        Name: "GitHub",        AppId: "259125845563242502",        Secret: "Bpw_g9c2TGXxfgLshDbSaL8tsCcqgczQ",        Enabled: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AuthProvider
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Key != model.Key {
        t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.AppId != model.AppId {
        t.Errorf("Expected AppId %v, got %v", model.AppId, result.AppId)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }}
