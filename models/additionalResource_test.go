package models

import (
    "encoding/json"
    "testing"
)

func TestAdditionalResourceModel(t *testing.T) {
    model := AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AdditionalResource
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Unit != model.Unit {
        t.Errorf("Expected Unit %v, got %v", model.Unit, result.Unit)
    }
    if result.Currency != model.Currency {
        t.Errorf("Expected Currency %v, got %v", model.Currency, result.Currency)
    }
    if result.Price != model.Price {
        t.Errorf("Expected Price %v, got %v", model.Price, result.Price)
    }
    if result.Value != model.Value {
        t.Errorf("Expected Value %v, got %v", model.Value, result.Value)
    }
    if result.InvoiceDesc != model.InvoiceDesc {
        t.Errorf("Expected InvoiceDesc %v, got %v", model.InvoiceDesc, result.InvoiceDesc)
    }}
