package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2NotionModel(t *testing.T) {
    model := OAuth2Notion{        Id: "github",        Enabled: true,        OauthClientId: "341d8700-0000-0000-0000-000000446ee3",        OauthClientSecret: "&lt;CLIENT_SECRET&gt;",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Notion
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
    if result.OauthClientId != model.OauthClientId {
        t.Errorf("Expected OauthClientId %v, got %v", model.OauthClientId, result.OauthClientId)
    }
    if result.OauthClientSecret != model.OauthClientSecret {
        t.Errorf("Expected OauthClientSecret %v, got %v", model.OauthClientSecret, result.OauthClientSecret)
    }}
