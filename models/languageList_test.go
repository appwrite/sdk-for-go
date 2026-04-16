package models

import (
    "encoding/json"
    "testing"
)

func TestLanguageListModel(t *testing.T) {
    model := LanguageList{        Total: 5,        Languages: []Language{Language{        Name: "Italian",        Code: "it",        NativeName: "Italiano",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result LanguageList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
