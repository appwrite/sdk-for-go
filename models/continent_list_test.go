package models

import (
    "encoding/json"
    "testing"
)

func TestContinentListModel(t *testing.T) {
    model := ContinentList{        Total: 5,        Continents: []Continent{Continent{        Name: "Europe",        Code: "EU",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ContinentList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
