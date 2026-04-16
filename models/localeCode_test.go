package models

import (
    "encoding/json"
    "testing"
)

func TestLocaleCodeModel(t *testing.T) {
    model := LocaleCode{        Code: "en-us",        Name: "US",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result LocaleCode
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Code != model.Code {
        t.Errorf("Expected Code %v, got %v", model.Code, result.Code)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }}
