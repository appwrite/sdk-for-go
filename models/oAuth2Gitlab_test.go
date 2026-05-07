package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2GitlabModel(t *testing.T) {
    model := OAuth2Gitlab{        Id: "github",        Enabled: true,        ApplicationId: "d41ffe0000000000000000000000000000000000000000000000000000d5e252",        Secret: "gloas-838cfa0000000000000000000000000000000000000000000000000000ecbb38",        Endpoint: "https://gitlab.com",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Gitlab
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
    if result.ApplicationId != model.ApplicationId {
        t.Errorf("Expected ApplicationId %v, got %v", model.ApplicationId, result.ApplicationId)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }
    if result.Endpoint != model.Endpoint {
        t.Errorf("Expected Endpoint %v, got %v", model.Endpoint, result.Endpoint)
    }}
