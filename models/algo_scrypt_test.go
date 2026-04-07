package models

import (
    "encoding/json"
    "testing"
)

func TestAlgoScryptModel(t *testing.T) {
    model := AlgoScrypt{        Type: "scrypt",        CostCpu: 8,        CostMemory: 14,        CostParallel: 1,        Length: 64,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AlgoScrypt
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.CostCpu != model.CostCpu {
        t.Errorf("Expected CostCpu %v, got %v", model.CostCpu, result.CostCpu)
    }
    if result.CostMemory != model.CostMemory {
        t.Errorf("Expected CostMemory %v, got %v", model.CostMemory, result.CostMemory)
    }
    if result.CostParallel != model.CostParallel {
        t.Errorf("Expected CostParallel %v, got %v", model.CostParallel, result.CostParallel)
    }
    if result.Length != model.Length {
        t.Errorf("Expected Length %v, got %v", model.Length, result.Length)
    }}
