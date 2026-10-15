package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDomainsListModel(t *testing.T) {
	model := DomainsList{Total: 5, Domains: []Domain{Domain{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Domain: "example.com", Registrar: "appwrite", Nameservers: "Appwrite", Expire: "2020-10-15T06:38:00.000+00:00", Renewal: "2020-10-15T06:38:00.000+00:00", AutoRenewal: true, RenewalPrice: 2599, TeamId: "5e5ea5c16897e", DnsRecords: []DnsRecord{DnsRecord{Id: "5f40a6e10c65e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Type: "A", Name: "mail", Value: "192.0.2.1", Ttl: 86400, Priority: 10, Lock: true, Weight: 10, Port: 443, Comment: "Mail server record"}}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DomainsList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestDomainsListModelDefaults(t *testing.T) {
	var fresh DomainsList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := DomainsList{Total: 5, Domains: []Domain{Domain{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Domain: "example.com", Registrar: "appwrite", Nameservers: "Appwrite", Expire: "2020-10-15T06:38:00.000+00:00", Renewal: "2020-10-15T06:38:00.000+00:00", AutoRenewal: true, RenewalPrice: 2599, TeamId: "5e5ea5c16897e", DnsRecords: []DnsRecord{DnsRecord{Id: "5f40a6e10c65e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Type: "A", Name: "mail", Value: "192.0.2.1", Ttl: 86400, Priority: 10, Lock: true, Weight: 10, Port: 443, Comment: "Mail server record"}}}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
