package models

import (
    "encoding/json"
    "testing"
)

func TestPolicyMfaFactorsModel(t *testing.T) {
    model := PolicyMfaFactors{        Id: "password-dictionary",        Totp: true,        Email: true,        Phone: true,        Custom: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PolicyMfaFactors
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Totp != model.Totp {
        t.Errorf("Expected Totp %v, got %v", model.Totp, result.Totp)
    }
    if result.Email != model.Email {
        t.Errorf("Expected Email %v, got %v", model.Email, result.Email)
    }
    if result.Phone != model.Phone {
        t.Errorf("Expected Phone %v, got %v", model.Phone, result.Phone)
    }
    if result.Custom != model.Custom {
        t.Errorf("Expected Custom %v, got %v", model.Custom, result.Custom)
    }}
