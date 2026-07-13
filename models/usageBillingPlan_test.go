package models

import (
    "encoding/json"
    "testing"
)

func TestUsageBillingPlanModel(t *testing.T) {
    model := UsageBillingPlan{        Bandwidth: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        Executions: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        Member: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        Realtime: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        RealtimeMessages: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        RealtimeBandwidth: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        Storage: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        Users: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        GBHours: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        ImageTransformations: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },        Credits: AdditionalResource{        Name: "string",        Unit: "GB",        Currency: "USD",        Price: 5,        Value: 25,        InvoiceDesc: "string",    },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result UsageBillingPlan
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }}
