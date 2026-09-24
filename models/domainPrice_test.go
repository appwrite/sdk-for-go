package models

import (
	"encoding/json"
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
