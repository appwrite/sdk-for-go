package models

import (
    "encoding/json"
    "testing"
)

func TestAlgoShaModel(t *testing.T) {
    model := AlgoSha{        Type: "sha",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AlgoSha
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }}
