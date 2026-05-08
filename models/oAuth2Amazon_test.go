package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2AmazonModel(t *testing.T) {
    model := OAuth2Amazon{        Id: "github",        Enabled: true,        ClientId: "amzn1.application-oa2-client.87400c00000000000000000000063d5b2",        ClientSecret: "&lt;CLIENT_SECRET&gt;",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Amazon
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
    }}
