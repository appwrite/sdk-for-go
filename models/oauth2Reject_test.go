package models

import (
    "encoding/json"
    "testing"
)

func TestOauth2RejectModel(t *testing.T) {
    model := Oauth2Reject{        RedirectUrl: "https://example.com/callback?error=access_denied&state=fghij",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Oauth2Reject
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.RedirectUrl != model.RedirectUrl {
        t.Errorf("Expected RedirectUrl %v, got %v", model.RedirectUrl, result.RedirectUrl)
    }}
