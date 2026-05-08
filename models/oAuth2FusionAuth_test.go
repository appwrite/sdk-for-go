package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2FusionAuthModel(t *testing.T) {
    model := OAuth2FusionAuth{        Id: "github",        Enabled: true,        ClientId: "b2222c00-0000-0000-0000-000000862097",        ClientSecret: "Jx4s0C0000000000000000000000000000000wGqLsc",        Endpoint: "example.fusionauth.io",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2FusionAuth
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
