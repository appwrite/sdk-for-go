package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2StripeModel(t *testing.T) {
    model := OAuth2Stripe{        Id: "github",        Enabled: true,        ClientId: "ca_UKibXX0000000000000000000006byvR",        ApiSecretKey: "sk_51SfOd000000000000000000000000000000000000000000000000000000000000000000000000000000000000000QGWYfp",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Stripe
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
    if result.ApiSecretKey != model.ApiSecretKey {
        t.Errorf("Expected ApiSecretKey %v, got %v", model.ApiSecretKey, result.ApiSecretKey)
    }}
