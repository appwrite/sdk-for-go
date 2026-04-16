package models

import (
    "encoding/json"
    "testing"
)

func TestAlgoScryptModifiedModel(t *testing.T) {
    model := AlgoScryptModified{        Type: "scryptMod",        Salt: "UxLMreBr6tYyjQ==",        SaltSeparator: "Bw==",        SignerKey: "XyEKE9RcTDeLEsL/RjwPDBv/RqDl8fb3gpYEOQaPihbxf1ZAtSOHCjuAAa7Q3oHpCYhXSN9tizHgVOwn6krflQ==",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AlgoScryptModified
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.Salt != model.Salt {
        t.Errorf("Expected Salt %v, got %v", model.Salt, result.Salt)
    }
    if result.SaltSeparator != model.SaltSeparator {
        t.Errorf("Expected SaltSeparator %v, got %v", model.SaltSeparator, result.SaltSeparator)
    }
    if result.SignerKey != model.SignerKey {
        t.Errorf("Expected SignerKey %v, got %v", model.SignerKey, result.SignerKey)
    }}
