package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2DailymotionModel(t *testing.T) {
    model := OAuth2Dailymotion{        Id: "github",        Enabled: true,        ApiKey: "07a9000000000000067f",        ApiSecret: "a399a90000000000000000000000000000d90639",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Dailymotion
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
    if result.ApiKey != model.ApiKey {
        t.Errorf("Expected ApiKey %v, got %v", model.ApiKey, result.ApiKey)
    }
    if result.ApiSecret != model.ApiSecret {
        t.Errorf("Expected ApiSecret %v, got %v", model.ApiSecret, result.ApiSecret)
    }}
