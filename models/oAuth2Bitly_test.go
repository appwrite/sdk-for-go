package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2BitlyModel(t *testing.T) {
    model := OAuth2Bitly{        Id: "github",        Enabled: true,        ClientId: "d95151000000000000000000000000000067af9b",        ClientSecret: "a13e250000000000000000000000000000d73095",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Bitly
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
    if result.ClientId != model.ClientId {
        t.Errorf("Expected ClientId %v, got %v", model.ClientId, result.ClientId)
    }
    if result.ClientSecret != model.ClientSecret {
        t.Errorf("Expected ClientSecret %v, got %v", model.ClientSecret, result.ClientSecret)
    }}
