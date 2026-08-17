package models

import (
	"encoding/json"
	"testing"
)

func TestBillingLimitsModel(t *testing.T) {
	model := BillingLimits{}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result BillingLimits
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
}
