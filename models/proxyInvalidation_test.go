package models

import (
    "encoding/json"
    "testing"
)

func TestProxyInvalidationModel(t *testing.T) {
    model := ProxyInvalidation{        Domain: "appwrite.company.com",        Type: "tag",        Reference: "products",        Status: "success",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ProxyInvalidation
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Domain != model.Domain {
        t.Errorf("Expected Domain %v, got %v", model.Domain, result.Domain)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.Reference != model.Reference {
        t.Errorf("Expected Reference %v, got %v", model.Reference, result.Reference)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }}
