package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2OidcModel(t *testing.T) {
    model := OAuth2Oidc{        Id: "github",        Enabled: true,        ClientId: "qibI2x0000000000000000000000000006L2YFoG",        ClientSecret: "Ah68ed000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003qpcHV",        WellKnownURL: "https://myoauth.com/.well-known/openid-configuration",        AuthorizationURL: "https://myoauth.com/oauth2/authorize",        TokenURL: "https://myoauth.com/oauth2/token",        UserInfoURL: "https://myoauth.com/oauth2/userinfo",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Oidc
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
    if result.WellKnownURL != model.WellKnownURL {
        t.Errorf("Expected WellKnownURL %v, got %v", model.WellKnownURL, result.WellKnownURL)
    }
    if result.AuthorizationURL != model.AuthorizationURL {
        t.Errorf("Expected AuthorizationURL %v, got %v", model.AuthorizationURL, result.AuthorizationURL)
    }
    if result.TokenURL != model.TokenURL {
        t.Errorf("Expected TokenURL %v, got %v", model.TokenURL, result.TokenURL)
    }
    if result.UserInfoURL != model.UserInfoURL {
        t.Errorf("Expected UserInfoURL %v, got %v", model.UserInfoURL, result.UserInfoURL)
    }}
