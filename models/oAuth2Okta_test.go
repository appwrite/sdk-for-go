package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2OktaModel(t *testing.T) {
    model := OAuth2Okta{        Id: "github",        Enabled: true,        ClientId: "0oa00000000000000698",        ClientSecret: "Kiq0000000000000000000000000000000000000-00000000000H2L5-3SJ-vRV",        Domain: "trial-6400025.okta.com",        AuthorizationServerId: "aus000000000000000h7z",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Okta
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
    if result.Domain != model.Domain {
        t.Errorf("Expected Domain %v, got %v", model.Domain, result.Domain)
    }
    if result.AuthorizationServerId != model.AuthorizationServerId {
        t.Errorf("Expected AuthorizationServerId %v, got %v", model.AuthorizationServerId, result.AuthorizationServerId)
    }}
