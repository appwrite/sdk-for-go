package models

import (
    "encoding/json"
    "testing"
)

func TestPhoneModel(t *testing.T) {
    model := Phone{        Code: "+1",        CountryCode: "US",        CountryName: "United States",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Phone
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Code != model.Code {
        t.Errorf("Expected Code %v, got %v", model.Code, result.Code)
    }
    if result.CountryCode != model.CountryCode {
        t.Errorf("Expected CountryCode %v, got %v", model.CountryCode, result.CountryCode)
    }
    if result.CountryName != model.CountryName {
        t.Errorf("Expected CountryName %v, got %v", model.CountryName, result.CountryName)
    }}
