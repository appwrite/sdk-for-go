package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2MicrosoftModel(t *testing.T) {
    model := OAuth2Microsoft{        Id: "github",        Enabled: true,        ApplicationId: "00001111-aaaa-2222-bbbb-3333cccc4444",        ApplicationSecret: "&lt;CLIENT_SECRET&gt;",        Tenant: "common",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Microsoft
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
    if result.ApplicationId != model.ApplicationId {
        t.Errorf("Expected ApplicationId %v, got %v", model.ApplicationId, result.ApplicationId)
    }
    if result.ApplicationSecret != model.ApplicationSecret {
        t.Errorf("Expected ApplicationSecret %v, got %v", model.ApplicationSecret, result.ApplicationSecret)
    }
    if result.Tenant != model.Tenant {
        t.Errorf("Expected Tenant %v, got %v", model.Tenant, result.Tenant)
    }}
