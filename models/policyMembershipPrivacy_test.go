package models

import (
    "encoding/json"
    "testing"
)

func TestPolicyMembershipPrivacyModel(t *testing.T) {
    model := PolicyMembershipPrivacy{        Id: "password-dictionary",        UserId: true,        UserEmail: true,        UserPhone: true,        UserName: true,        UserMFA: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PolicyMembershipPrivacy
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.UserEmail != model.UserEmail {
        t.Errorf("Expected UserEmail %v, got %v", model.UserEmail, result.UserEmail)
    }
    if result.UserPhone != model.UserPhone {
        t.Errorf("Expected UserPhone %v, got %v", model.UserPhone, result.UserPhone)
    }
    if result.UserName != model.UserName {
        t.Errorf("Expected UserName %v, got %v", model.UserName, result.UserName)
    }
    if result.UserMFA != model.UserMFA {
        t.Errorf("Expected UserMFA %v, got %v", model.UserMFA, result.UserMFA)
    }}
