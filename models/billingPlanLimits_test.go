package models

import (
    "encoding/json"
    "testing"
)

func TestBillingPlanLimitsModel(t *testing.T) {
    model := BillingPlanLimits{    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BillingPlanLimits
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }}
