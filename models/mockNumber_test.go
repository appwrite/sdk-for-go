package models

import (
    "encoding/json"
    "testing"
)

func TestMockNumberModel(t *testing.T) {
    model := MockNumber{        Phone: "+1612842323",        Otp: "123456",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result MockNumber
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Phone != model.Phone {
        t.Errorf("Expected Phone %v, got %v", model.Phone, result.Phone)
    }
    if result.Otp != model.Otp {
        t.Errorf("Expected Otp %v, got %v", model.Otp, result.Otp)
    }}
