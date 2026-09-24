package models

import (
	"encoding/json"
	"testing"
)

func TestDomainModel(t *testing.T) {
	model := Domain{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Domain: "example.com", Registrar: "appwrite", Nameservers: "Appwrite", Expire: "2020-10-15T06:38:00.000+00:00", Renewal: "2020-10-15T06:38:00.000+00:00", AutoRenewal: true, RenewalPrice: 2599, TeamId: "5e5ea5c16897e", DnsRecords: []DnsRecord{DnsRecord{Id: "5f40a6e10c65e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Type: "A", Name: "mail", Value: "192.0.2.1", Ttl: 86400, Priority: 10, Lock: true, Weight: 10, Port: 443, Comment: "Mail server record"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Domain
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
	if result.Domain != model.Domain {
		t.Errorf("Expected Domain %v, got %v", model.Domain, result.Domain)
	}
	if result.Registrar != model.Registrar {
		t.Errorf("Expected Registrar %v, got %v", model.Registrar, result.Registrar)
	}
	if result.Nameservers != model.Nameservers {
		t.Errorf("Expected Nameservers %v, got %v", model.Nameservers, result.Nameservers)
	}
	if result.Expire != model.Expire {
		t.Errorf("Expected Expire %v, got %v", model.Expire, result.Expire)
	}
	if result.Renewal != model.Renewal {
		t.Errorf("Expected Renewal %v, got %v", model.Renewal, result.Renewal)
	}
	if result.AutoRenewal != model.AutoRenewal {
		t.Errorf("Expected AutoRenewal %v, got %v", model.AutoRenewal, result.AutoRenewal)
	}
	if result.RenewalPrice != model.RenewalPrice {
		t.Errorf("Expected RenewalPrice %v, got %v", model.RenewalPrice, result.RenewalPrice)
	}
	if result.TeamId != model.TeamId {
		t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
	}
}
