package models

import (
    "encoding/json"
    "testing"
)

func TestAlgoArgon2Model(t *testing.T) {
    model := AlgoArgon2{        Type: "argon2",        MemoryCost: 65536,        TimeCost: 4,        Threads: 3,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AlgoArgon2
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.MemoryCost != model.MemoryCost {
        t.Errorf("Expected MemoryCost %v, got %v", model.MemoryCost, result.MemoryCost)
    }
    if result.TimeCost != model.TimeCost {
        t.Errorf("Expected TimeCost %v, got %v", model.TimeCost, result.TimeCost)
    }
    if result.Threads != model.Threads {
        t.Errorf("Expected Threads %v, got %v", model.Threads, result.Threads)
    }}
