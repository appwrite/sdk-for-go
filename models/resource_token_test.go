package models

import (
    "encoding/json"
    "testing"
)

func TestResourceTokenModel(t *testing.T) {
    model := ResourceToken{        Id: "bb8ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        ResourceId: "5e5ea5c168bb8:5e5ea5c168bb8",        ResourceType: "files",        Expire: "2020-10-15T06:38:00.000+00:00",        Secret: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",        AccessedAt: "2020-10-15T06:38:00.000+00:00",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ResourceToken
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
    if result.ResourceId != model.ResourceId {
        t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
    }
    if result.ResourceType != model.ResourceType {
        t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
    }
    if result.Expire != model.Expire {
        t.Errorf("Expected Expire %v, got %v", model.Expire, result.Expire)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }
    if result.AccessedAt != model.AccessedAt {
        t.Errorf("Expected AccessedAt %v, got %v", model.AccessedAt, result.AccessedAt)
    }}
