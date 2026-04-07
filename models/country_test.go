package models

import (
    "encoding/json"
    "testing"
)

func TestCountryModel(t *testing.T) {
    model := Country{        Name: "United States",        Code: "US",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Country
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Code != model.Code {
        t.Errorf("Expected Code %v, got %v", model.Code, result.Code)
    }}
