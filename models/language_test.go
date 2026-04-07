package models

import (
    "encoding/json"
    "testing"
)

func TestLanguageModel(t *testing.T) {
    model := Language{        Name: "Italian",        Code: "it",        NativeName: "Italiano",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Language
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Code != model.Code {
        t.Errorf("Expected Code %v, got %v", model.Code, result.Code)
    }
    if result.NativeName != model.NativeName {
        t.Errorf("Expected NativeName %v, got %v", model.NativeName, result.NativeName)
    }}
