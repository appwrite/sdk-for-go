package models

import (
    "encoding/json"
    "testing"
)

func TestHeadersModel(t *testing.T) {
    model := Headers{        Name: "Content-Type",        Value: "application/json",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Headers
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Value != model.Value {
        t.Errorf("Expected Value %v, got %v", model.Value, result.Value)
    }}
