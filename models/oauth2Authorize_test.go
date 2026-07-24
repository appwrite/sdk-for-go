package models

import (
    "encoding/json"
    "testing"
)

func TestOauth2AuthorizeModel(t *testing.T) {
    model := Oauth2Authorize{        GrantId: "5e5ea5c16897e",        RedirectUrl: "https://example.com/callback?code=abcde&state=fghij",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Oauth2Authorize
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.GrantId != model.GrantId {
        t.Errorf("Expected GrantId %v, got %v", model.GrantId, result.GrantId)
    }
    if result.RedirectUrl != model.RedirectUrl {
        t.Errorf("Expected RedirectUrl %v, got %v", model.RedirectUrl, result.RedirectUrl)
    }}
