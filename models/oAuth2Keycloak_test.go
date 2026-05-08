package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2KeycloakModel(t *testing.T) {
    model := OAuth2Keycloak{        Id: "github",        Enabled: true,        ClientId: "appwrite-o0000000st-app",        ClientSecret: "&lt;CLIENT_SECRET&gt;",        Endpoint: "keycloak.example.com",        RealmName: "appwrite-realm",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Keycloak
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
    }
    if result.RealmName != model.RealmName {
        t.Errorf("Expected RealmName %v, got %v", model.RealmName, result.RealmName)
    }}
