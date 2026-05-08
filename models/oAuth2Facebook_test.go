package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2FacebookModel(t *testing.T) {
    model := OAuth2Facebook{        Id: "github",        Enabled: true,        AppId: "260600000007694",        AppSecret: "2d0b2800000000000000000000d38af4",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Facebook
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
    if result.AppId != model.AppId {
        t.Errorf("Expected AppId %v, got %v", model.AppId, result.AppId)
    }
    if result.AppSecret != model.AppSecret {
        t.Errorf("Expected AppSecret %v, got %v", model.AppSecret, result.AppSecret)
    }}
