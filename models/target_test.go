package models

import (
    "encoding/json"
    "testing"
)

func TestTargetModel(t *testing.T) {
    model := Target{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Apple iPhone 12",        UserId: "259125845563242502",        ProviderType: "email",        Identifier: "token",        Expired: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Target
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.ProviderType != model.ProviderType {
        t.Errorf("Expected ProviderType %v, got %v", model.ProviderType, result.ProviderType)
    }
    if result.Identifier != model.Identifier {
        t.Errorf("Expected Identifier %v, got %v", model.Identifier, result.Identifier)
    }
    if result.Expired != model.Expired {
        t.Errorf("Expected Expired %v, got %v", model.Expired, result.Expired)
    }}
