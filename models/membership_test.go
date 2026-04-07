package models

import (
    "encoding/json"
    "testing"
)

func TestMembershipModel(t *testing.T) {
    model := Membership{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        UserId: "5e5ea5c16897e",        UserName: "John Doe",        UserEmail: "john@appwrite.io",        TeamId: "5e5ea5c16897e",        TeamName: "VIP",        Invited: "2020-10-15T06:38:00.000+00:00",        Joined: "2020-10-15T06:38:00.000+00:00",        Confirm: true,        Mfa: true,        Roles: []string{"test"},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Membership
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
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.UserName != model.UserName {
        t.Errorf("Expected UserName %v, got %v", model.UserName, result.UserName)
    }
    if result.UserEmail != model.UserEmail {
        t.Errorf("Expected UserEmail %v, got %v", model.UserEmail, result.UserEmail)
    }
    if result.TeamId != model.TeamId {
        t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
    }
    if result.TeamName != model.TeamName {
        t.Errorf("Expected TeamName %v, got %v", model.TeamName, result.TeamName)
    }
    if result.Invited != model.Invited {
        t.Errorf("Expected Invited %v, got %v", model.Invited, result.Invited)
    }
    if result.Joined != model.Joined {
        t.Errorf("Expected Joined %v, got %v", model.Joined, result.Joined)
    }
    if result.Confirm != model.Confirm {
        t.Errorf("Expected Confirm %v, got %v", model.Confirm, result.Confirm)
    }
    if result.Mfa != model.Mfa {
        t.Errorf("Expected Mfa %v, got %v", model.Mfa, result.Mfa)
    }}
