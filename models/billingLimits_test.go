package models

import (
    "encoding/json"
    "testing"
)

func TestBillingLimitsModel(t *testing.T) {
    model := BillingLimits{        Bandwidth: 5,        Storage: 150,        Users: 200000,        Executions: 750000,        GBHours: 100,        ImageTransformations: 100,        AuthPhone: 10,        BudgetLimit: 100,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BillingLimits
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Bandwidth != model.Bandwidth {
        t.Errorf("Expected Bandwidth %v, got %v", model.Bandwidth, result.Bandwidth)
    }
    if result.Storage != model.Storage {
        t.Errorf("Expected Storage %v, got %v", model.Storage, result.Storage)
    }
    if result.Users != model.Users {
        t.Errorf("Expected Users %v, got %v", model.Users, result.Users)
    }
    if result.Executions != model.Executions {
        t.Errorf("Expected Executions %v, got %v", model.Executions, result.Executions)
    }
    if result.GBHours != model.GBHours {
        t.Errorf("Expected GBHours %v, got %v", model.GBHours, result.GBHours)
    }
    if result.ImageTransformations != model.ImageTransformations {
        t.Errorf("Expected ImageTransformations %v, got %v", model.ImageTransformations, result.ImageTransformations)
    }
    if result.AuthPhone != model.AuthPhone {
        t.Errorf("Expected AuthPhone %v, got %v", model.AuthPhone, result.AuthPhone)
    }
    if result.BudgetLimit != model.BudgetLimit {
        t.Errorf("Expected BudgetLimit %v, got %v", model.BudgetLimit, result.BudgetLimit)
    }}
