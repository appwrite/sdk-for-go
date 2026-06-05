package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2DisqusModel(t *testing.T) {
    model := OAuth2Disqus{        Id: "github",        Enabled: true,        PublicKey: "cgegH70000000000000000000000000000000000000000000000000000Hr1nYX",        SecretKey: "W7Bykj00000000000000000000000000000000000000000000000000003o43w9",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Disqus
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
    if result.PublicKey != model.PublicKey {
        t.Errorf("Expected PublicKey %v, got %v", model.PublicKey, result.PublicKey)
    }
    if result.SecretKey != model.SecretKey {
        t.Errorf("Expected SecretKey %v, got %v", model.SecretKey, result.SecretKey)
    }}
