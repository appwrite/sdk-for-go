package models

import (
	"encoding/json"
	"testing"
)

func TestBillingPlanAddonModel(t *testing.T) {
	model := BillingPlanAddon{}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result BillingPlanAddon
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
}
