package models

import (
    "encoding/json"
    "testing"
)

func TestCurrencyModel(t *testing.T) {
    model := Currency{        Symbol: "$",        Name: "US dollar",        SymbolNative: "$",        DecimalDigits: 2,        Rounding: 0,        Code: "USD",        NamePlural: "US dollars",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Currency
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Symbol != model.Symbol {
        t.Errorf("Expected Symbol %v, got %v", model.Symbol, result.Symbol)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.SymbolNative != model.SymbolNative {
        t.Errorf("Expected SymbolNative %v, got %v", model.SymbolNative, result.SymbolNative)
    }
    if result.DecimalDigits != model.DecimalDigits {
        t.Errorf("Expected DecimalDigits %v, got %v", model.DecimalDigits, result.DecimalDigits)
    }
    if result.Rounding != model.Rounding {
        t.Errorf("Expected Rounding %v, got %v", model.Rounding, result.Rounding)
    }
    if result.Code != model.Code {
        t.Errorf("Expected Code %v, got %v", model.Code, result.Code)
    }
    if result.NamePlural != model.NamePlural {
        t.Errorf("Expected NamePlural %v, got %v", model.NamePlural, result.NamePlural)
    }}
