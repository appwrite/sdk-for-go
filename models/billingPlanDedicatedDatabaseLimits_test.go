package models

import (
    "encoding/json"
    "testing"
)

func TestBillingPlanDedicatedDatabaseLimitsModel(t *testing.T) {
    model := BillingPlanDedicatedDatabaseLimits{    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BillingPlanDedicatedDatabaseLimits
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }}
