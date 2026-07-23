package models

import (
    "encoding/json"
    "testing"
)

func TestOauth2GrantModel(t *testing.T) {
    model := Oauth2Grant{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        UserId: "5e5ea5c16897e",        AppId: "5e5ea5c16897e",        Scopes: []string{"test"},        Resources: []string{"test"},        AuthorizationDetails: "[{\"type\":\"calendar\",\"identifier\":\"primary\",\"actions\":[\"read_events\",\"create_event\"]}]",        Prompt: "login",        RedirectUri: "https://example.com/callback",        AuthTime: 1592981250,        Expire: "2020-10-15T06:38:00.000+00:00",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Oauth2Grant
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.AppId != model.AppId {
        t.Errorf("Expected AppId %v, got %v", model.AppId, result.AppId)
    }
    if result.AuthorizationDetails != model.AuthorizationDetails {
        t.Errorf("Expected AuthorizationDetails %v, got %v", model.AuthorizationDetails, result.AuthorizationDetails)
    }
    if result.Prompt != model.Prompt {
        t.Errorf("Expected Prompt %v, got %v", model.Prompt, result.Prompt)
    }
    if result.RedirectUri != model.RedirectUri {
        t.Errorf("Expected RedirectUri %v, got %v", model.RedirectUri, result.RedirectUri)
    }
    if result.AuthTime != model.AuthTime {
        t.Errorf("Expected AuthTime %v, got %v", model.AuthTime, result.AuthTime)
    }
    if result.Expire != model.Expire {
        t.Errorf("Expected Expire %v, got %v", model.Expire, result.Expire)
    }}
