package models

import (
	"encoding/json"
	"testing"
)

func TestDnsRecordModel(t *testing.T) {
	model := DnsRecord{Id: "5f40a6e10c65e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Type: "A", Name: "mail", Value: "192.0.2.1", Ttl: 86400, Priority: 10, Lock: true, Weight: 10, Port: 443, Comment: "Mail server record"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DnsRecord
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
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Value != model.Value {
		t.Errorf("Expected Value %v, got %v", model.Value, result.Value)
	}
	if result.Ttl != model.Ttl {
		t.Errorf("Expected Ttl %v, got %v", model.Ttl, result.Ttl)
	}
	if result.Priority != model.Priority {
		t.Errorf("Expected Priority %v, got %v", model.Priority, result.Priority)
	}
	if result.Lock != model.Lock {
		t.Errorf("Expected Lock %v, got %v", model.Lock, result.Lock)
	}
	if result.Weight != model.Weight {
		t.Errorf("Expected Weight %v, got %v", model.Weight, result.Weight)
	}
	if result.Port != model.Port {
		t.Errorf("Expected Port %v, got %v", model.Port, result.Port)
	}
	if result.Comment != model.Comment {
		t.Errorf("Expected Comment %v, got %v", model.Comment, result.Comment)
	}
}
