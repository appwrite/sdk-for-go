package models

import (
    "encoding/json"
    "testing"
)

func TestCountryListModel(t *testing.T) {
    model := CountryList{        Total: 5,        Countries: []Country{Country{        Name: "United States",        Code: "US",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result CountryList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
