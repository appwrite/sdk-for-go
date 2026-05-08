package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2Auth0Model(t *testing.T) {
    model := OAuth2Auth0{        Id: "github",        Enabled: true,        ClientId: "OaOkIA000000000000000000005KLSYq",        ClientSecret: "&lt;CLIENT_SECRET&gt;",        Endpoint: "example.us.auth0.com",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Auth0
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
