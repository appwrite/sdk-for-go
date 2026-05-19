package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2BitbucketModel(t *testing.T) {
    model := OAuth2Bitbucket{        Id: "github",        Enabled: true,        Key: "Knt70000000000ByRc",        Secret: "NMfLZJ00000000000000000000TLQdDx",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Bitbucket
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
    if result.Key != model.Key {
        t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }}
