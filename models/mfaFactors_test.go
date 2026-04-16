package models

import (
    "encoding/json"
    "testing"
)

func TestMfaFactorsModel(t *testing.T) {
    model := MfaFactors{        Totp: true,        Phone: true,        Email: true,        RecoveryCode: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result MfaFactors
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Totp != model.Totp {
        t.Errorf("Expected Totp %v, got %v", model.Totp, result.Totp)
    }
    if result.Phone != model.Phone {
        t.Errorf("Expected Phone %v, got %v", model.Phone, result.Phone)
    }
    if result.Email != model.Email {
        t.Errorf("Expected Email %v, got %v", model.Email, result.Email)
    }
    if result.RecoveryCode != model.RecoveryCode {
        t.Errorf("Expected RecoveryCode %v, got %v", model.RecoveryCode, result.RecoveryCode)
    }}
