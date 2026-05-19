package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2PaypalModel(t *testing.T) {
    model := OAuth2Paypal{        Id: "github",        Enabled: true,        ClientId: "AdhIEG7-000000000000-0000000000000000000000000000000-0000000000000000000000-2pyB",        SecretKey: "EH8KCXtew--000000000000000000000000000000000000000_C-1_5UP_000000000000000CB7KDp",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Paypal
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
    if result.SecretKey != model.SecretKey {
        t.Errorf("Expected SecretKey %v, got %v", model.SecretKey, result.SecretKey)
    }}
