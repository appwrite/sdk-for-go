package models

import (
    "encoding/json"
    "testing"
)

func TestAlgoBcryptModel(t *testing.T) {
    model := AlgoBcrypt{        Type: "bcrypt",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AlgoBcrypt
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }}
