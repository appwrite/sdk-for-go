package models

import (
    "encoding/json"
    "testing"
)

func TestUserModel(t *testing.T) {
    model := User{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "John Doe",        Registration: "2020-10-15T06:38:00.000+00:00",        Status: true,        Labels: []string{"test"},        PasswordUpdate: "2020-10-15T06:38:00.000+00:00",        Email: "john@appwrite.io",        Phone: "+4930901820",        EmailVerification: true,        PhoneVerification: true,        Mfa: true,        Prefs: Preferences{    },        Targets: []Target{Target{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Apple iPhone 12",        UserId: "259125845563242502",        ProviderType: "email",        Identifier: "token",        Expired: true,    },
            },        AccessedAt: "2020-10-15T06:38:00.000+00:00",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result User
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
    if result.Registration != model.Registration {
        t.Errorf("Expected Registration %v, got %v", model.Registration, result.Registration)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.PasswordUpdate != model.PasswordUpdate {
        t.Errorf("Expected PasswordUpdate %v, got %v", model.PasswordUpdate, result.PasswordUpdate)
    }
    if result.Email != model.Email {
        t.Errorf("Expected Email %v, got %v", model.Email, result.Email)
    }
    if result.Phone != model.Phone {
        t.Errorf("Expected Phone %v, got %v", model.Phone, result.Phone)
    }
    if result.EmailVerification != model.EmailVerification {
        t.Errorf("Expected EmailVerification %v, got %v", model.EmailVerification, result.EmailVerification)
    }
    if result.PhoneVerification != model.PhoneVerification {
        t.Errorf("Expected PhoneVerification %v, got %v", model.PhoneVerification, result.PhoneVerification)
    }
    if result.Mfa != model.Mfa {
        t.Errorf("Expected Mfa %v, got %v", model.Mfa, result.Mfa)
    }
    if result.AccessedAt != model.AccessedAt {
        t.Errorf("Expected AccessedAt %v, got %v", model.AccessedAt, result.AccessedAt)
    }}
