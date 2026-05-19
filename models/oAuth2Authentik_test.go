package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2AuthentikModel(t *testing.T) {
    model := OAuth2Authentik{        Id: "github",        Enabled: true,        ClientId: "dTKOPa0000000000000000000000000000e7G8hv",        ClientSecret: "ntQadq000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000Hp5WK",        Endpoint: "example.authentik.com",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Authentik
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
    }
    if result.Endpoint != model.Endpoint {
        t.Errorf("Expected Endpoint %v, got %v", model.Endpoint, result.Endpoint)
    }}
