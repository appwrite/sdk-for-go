package models

import (
    "encoding/json"
    "testing"
)

func TestCurrencyListModel(t *testing.T) {
    model := CurrencyList{        Total: 5,        Currencies: []Currency{Currency{        Symbol: "$",        Name: "US dollar",        SymbolNative: "$",        DecimalDigits: 2,        Rounding: 0,        Code: "USD",        NamePlural: "US dollars",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result CurrencyList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
