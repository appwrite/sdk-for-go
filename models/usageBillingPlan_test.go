package models

import (
	"encoding/json"
	"testing"
)

func TestUsageBillingPlanModel(t *testing.T) {
	model := UsageBillingPlan{}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result UsageBillingPlan
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
}
