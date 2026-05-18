package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2LinkedinModel(t *testing.T) {
    model := OAuth2Linkedin{        Id: "github",        Enabled: true,        ClientId: "770000000000dv",        PrimaryClientSecret: "example-linkedin-client-secret",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Linkedin
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
    if result.PrimaryClientSecret != model.PrimaryClientSecret {
        t.Errorf("Expected PrimaryClientSecret %v, got %v", model.PrimaryClientSecret, result.PrimaryClientSecret)
    }}
