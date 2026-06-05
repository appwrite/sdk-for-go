package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2TradeshiftModel(t *testing.T) {
    model := OAuth2Tradeshift{        Id: "github",        Enabled: true,        Oauth2ClientId: "appwrite-test-org.appwrite-test-app",        Oauth2ClientSecret: "7cb52700-0000-0000-0000-000000ca5b83",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Tradeshift
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
    if result.Oauth2ClientId != model.Oauth2ClientId {
        t.Errorf("Expected Oauth2ClientId %v, got %v", model.Oauth2ClientId, result.Oauth2ClientId)
    }
    if result.Oauth2ClientSecret != model.Oauth2ClientSecret {
        t.Errorf("Expected Oauth2ClientSecret %v, got %v", model.Oauth2ClientSecret, result.Oauth2ClientSecret)
    }}
