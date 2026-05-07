package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2EtsyModel(t *testing.T) {
    model := OAuth2Etsy{        Id: "github",        Enabled: true,        KeyString: "nsgzxh0000000000008j85a2",        SharedSecret: "tp000000ru",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Etsy
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }
    if result.KeyString != model.KeyString {
        t.Errorf("Expected KeyString %v, got %v", model.KeyString, result.KeyString)
    }
    if result.SharedSecret != model.SharedSecret {
        t.Errorf("Expected SharedSecret %v, got %v", model.SharedSecret, result.SharedSecret)
    }}
