package models

import (
	"encoding/json"
	"testing"
)

func TestDnsRecordsListModel(t *testing.T) {
	model := DnsRecordsList{Total: 5, DnsRecords: []DnsRecord{DnsRecord{Id: "5f40a6e10c65e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Type: "A", Name: "mail", Value: "192.0.2.1", Ttl: 86400, Priority: 10, Lock: true, Weight: 10, Port: 443, Comment: "Mail server record"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DnsRecordsList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
