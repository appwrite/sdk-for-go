package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2XModel(t *testing.T) {
    model := OAuth2X{        Id: "github",        Enabled: true,        CustomerKey: "slzZV0000000000000NFLaWT",        SecretKey: "&lt;CLIENT_SECRET&gt;",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2X
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
    if result.CustomerKey != model.CustomerKey {
        t.Errorf("Expected CustomerKey %v, got %v", model.CustomerKey, result.CustomerKey)
    }
    if result.SecretKey != model.SecretKey {
        t.Errorf("Expected SecretKey %v, got %v", model.SecretKey, result.SecretKey)
    }}
