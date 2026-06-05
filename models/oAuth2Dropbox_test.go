package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2DropboxModel(t *testing.T) {
    model := OAuth2Dropbox{        Id: "github",        Enabled: true,        AppKey: "jl000000000009t",        AppSecret: "g200000000000vw",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Dropbox
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
    if result.AppKey != model.AppKey {
        t.Errorf("Expected AppKey %v, got %v", model.AppKey, result.AppKey)
    }
    if result.AppSecret != model.AppSecret {
        t.Errorf("Expected AppSecret %v, got %v", model.AppSecret, result.AppSecret)
    }}
