package models

import (
    "encoding/json"
    "testing"
)

func TestOAuth2SalesforceModel(t *testing.T) {
    model := OAuth2Salesforce{        Id: "github",        Enabled: true,        CustomerKey: "3MVG9I0000000000000000000000000000000000000000000000000000000000000000000000000C5Aejq",        CustomerSecret: "3w000000000000e2",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result OAuth2Salesforce
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
    if result.CustomerKey != model.CustomerKey {
        t.Errorf("Expected CustomerKey %v, got %v", model.CustomerKey, result.CustomerKey)
    }
    if result.CustomerSecret != model.CustomerSecret {
        t.Errorf("Expected CustomerSecret %v, got %v", model.CustomerSecret, result.CustomerSecret)
    }}
