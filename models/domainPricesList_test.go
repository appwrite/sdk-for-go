package models

import (
	"encoding/json"
	"testing"
)

func TestDomainPricesListModel(t *testing.T) {
	model := DomainPricesList{Total: 5, Prices: []DomainPrice{DomainPrice{Domain: "example.com", Tld: "com", Available: true, PeriodYears: 1, Premium: true, RenewalPeriodYears: 1}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DomainPricesList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
