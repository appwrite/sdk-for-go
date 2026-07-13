package models

import (
    "encoding/json"
    "testing"
)

func TestBillingPlanAddonDetailsModel(t *testing.T) {
    model := BillingPlanAddonDetails{        Supported: true,        PlanIncluded: 1,        Limit: 5,        Type: "numeric",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BillingPlanAddonDetails
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Supported != model.Supported {
        t.Errorf("Expected Supported %v, got %v", model.Supported, result.Supported)
    }
    if result.PlanIncluded != model.PlanIncluded {
        t.Errorf("Expected PlanIncluded %v, got %v", model.PlanIncluded, result.PlanIncluded)
    }
    if result.Limit != model.Limit {
        t.Errorf("Expected Limit %v, got %v", model.Limit, result.Limit)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
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
