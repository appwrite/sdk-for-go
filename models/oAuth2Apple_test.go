package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2AppleModel(t *testing.T) {
    model := OAuth2Apple{        Id: "apple",        Enabled: true,        ServiceId: "ip.appwrite.app.web",        KeyId: "P4000000N8",        TeamId: "D4000000R6",        P8File: "-----BEGIN PRIVATE KEY-----MIGTAg...jy2Xbna-----END PRIVATE KEY-----",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Apple
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
    if result.ServiceId != model.ServiceId {
        t.Errorf("Expected ServiceId %v, got %v", model.ServiceId, result.ServiceId)
    }
    if result.KeyId != model.KeyId {
        t.Errorf("Expected KeyId %v, got %v", model.KeyId, result.KeyId)
    }
    if result.TeamId != model.TeamId {
        t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
    }
    if result.P8File != model.P8File {
        t.Errorf("Expected P8File %v, got %v", model.P8File, result.P8File)
    }}
