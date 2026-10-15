package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBillingPlanSupportedAddonsModel(t *testing.T) {
	model := BillingPlanSupportedAddons{Baa: true, PremiumGeoDB: true, PremiumGeoDBOrg: true}

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
	}
}

func TestBillingPlanSupportedAddonsModelDefaults(t *testing.T) {
	var fresh BillingPlanSupportedAddons
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Baa)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing baa to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PremiumGeoDB)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing premiumGeoDB to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PremiumGeoDBOrg)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing premiumGeoDBOrg to default to %v, got %v", want, got)
		}
	}

	reused := BillingPlanSupportedAddons{Baa: true, PremiumGeoDB: true, PremiumGeoDBOrg: true}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
