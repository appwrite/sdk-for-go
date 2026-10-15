package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDomainPriceModel(t *testing.T) {
	model := DomainPrice{Domain: "example.com", Tld: "com", Available: true, PeriodYears: 1, Premium: true, RenewalPeriodYears: 1}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DomainPrice
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Domain != model.Domain {
		t.Errorf("Expected Domain %v, got %v", model.Domain, result.Domain)
	}
	if result.Tld != model.Tld {
		t.Errorf("Expected Tld %v, got %v", model.Tld, result.Tld)
	}
	if result.Available != model.Available {
		t.Errorf("Expected Available %v, got %v", model.Available, result.Available)
	}
	if result.PeriodYears != model.PeriodYears {
		t.Errorf("Expected PeriodYears %v, got %v", model.PeriodYears, result.PeriodYears)
	}
	if result.Premium != model.Premium {
		t.Errorf("Expected Premium %v, got %v", model.Premium, result.Premium)
	}
	if result.RenewalPeriodYears != model.RenewalPeriodYears {
		t.Errorf("Expected RenewalPeriodYears %v, got %v", model.RenewalPeriodYears, result.RenewalPeriodYears)
	}
}

func TestDomainPriceModelDefaults(t *testing.T) {
	var fresh DomainPrice
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Domain)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing domain to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Tld)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing tld to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Available)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing available to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("1"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PeriodYears)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing periodYears to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Premium)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing premium to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("1"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RenewalPeriodYears)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing renewalPeriodYears to default to %v, got %v", want, got)
		}
	}

	reused := DomainPrice{Domain: "example.com", Tld: "com", Available: true, PeriodYears: 1, Premium: true, RenewalPeriodYears: 1}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
