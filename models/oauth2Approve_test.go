package models

import (
    "encoding/json"
    "testing"
)

func TestOauth2ApproveModel(t *testing.T) {
    model := Oauth2Approve{        RedirectUrl: "https://example.com/callback?code=abcde&state=fghij",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Oauth2Approve
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.RedirectUrl != model.RedirectUrl {
        t.Errorf("Expected RedirectUrl %v, got %v", model.RedirectUrl, result.RedirectUrl)
    }}
