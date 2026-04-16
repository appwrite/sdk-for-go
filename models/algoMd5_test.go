package models

import (
    "encoding/json"
    "testing"
)

func TestAlgoMd5Model(t *testing.T) {
    model := AlgoMd5{        Type: "md5",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AlgoMd5
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }}
