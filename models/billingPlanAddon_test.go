package models

import (
    "encoding/json"
    "testing"
)

func TestBillingPlanAddonModel(t *testing.T) {
    model := BillingPlanAddon{        Seats: BillingPlanAddonDetails{        Supported: true,        PlanIncluded: 1,        Limit: 5,        Type: "numeric",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        Projects: BillingPlanAddonDetails{        Supported: true,        PlanIncluded: 1,        Limit: 5,        Type: "numeric",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BillingPlanAddon
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }}
