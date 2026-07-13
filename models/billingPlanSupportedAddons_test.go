package models

import (
    "encoding/json"
    "testing"
)

func TestBillingPlanSupportedAddonsModel(t *testing.T) {
    model := BillingPlanSupportedAddons{        Baa: true,        PremiumGeoDB: true,        PremiumGeoDBOrg: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BillingPlanSupportedAddons
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Baa != model.Baa {
        t.Errorf("Expected Baa %v, got %v", model.Baa, result.Baa)
    }
    if result.PremiumGeoDB != model.PremiumGeoDB {
        t.Errorf("Expected PremiumGeoDB %v, got %v", model.PremiumGeoDB, result.PremiumGeoDB)
    }
    if result.PremiumGeoDBOrg != model.PremiumGeoDBOrg {
        t.Errorf("Expected PremiumGeoDBOrg %v, got %v", model.PremiumGeoDBOrg, result.PremiumGeoDBOrg)
    }}
