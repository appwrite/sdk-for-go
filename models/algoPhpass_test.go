package models

import (
    "encoding/json"
    "testing"
)

func TestAlgoPhpassModel(t *testing.T) {
    model := AlgoPhpass{        Type: "phpass",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AlgoPhpass
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }}
