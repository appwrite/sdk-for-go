package models

import (
    "encoding/json"
    "testing"
)

func TestPresenceModel(t *testing.T) {
    model := Presence{        Id: "5e5ea5c16897e",        Sequence: "1",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        UserInternalId: "1",        UserId: "674af8f3e12a5f9ac0be",        Source: "HTTP",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Presence
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Sequence != model.Sequence {
        t.Errorf("Expected Sequence %v, got %v", model.Sequence, result.Sequence)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.UserInternalId != model.UserInternalId {
        t.Errorf("Expected UserInternalId %v, got %v", model.UserInternalId, result.UserInternalId)
    }
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.Source != model.Source {
        t.Errorf("Expected Source %v, got %v", model.Source, result.Source)
    }}
