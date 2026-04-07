package models

import (
    "encoding/json"
    "testing"
)

func TestSubscriberModel(t *testing.T) {
    model := Subscriber{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        TargetId: "259125845563242502",        Target: Target{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Apple iPhone 12",        UserId: "259125845563242502",        ProviderType: "email",        Identifier: "token",        Expired: true,    },        UserId: "5e5ea5c16897e",        UserName: "Aegon Targaryen",        TopicId: "259125845563242502",        ProviderType: "email",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Subscriber
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
    if result.TargetId != model.TargetId {
        t.Errorf("Expected TargetId %v, got %v", model.TargetId, result.TargetId)
    }
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.UserName != model.UserName {
        t.Errorf("Expected UserName %v, got %v", model.UserName, result.UserName)
    }
    if result.TopicId != model.TopicId {
        t.Errorf("Expected TopicId %v, got %v", model.TopicId, result.TopicId)
    }
    if result.ProviderType != model.ProviderType {
        t.Errorf("Expected ProviderType %v, got %v", model.ProviderType, result.ProviderType)
    }}
