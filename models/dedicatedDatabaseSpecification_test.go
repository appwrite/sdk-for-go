package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseSpecificationModel(t *testing.T) {
	model := DedicatedDatabaseSpecification{Slug: "s-2vcpu-2gb", Name: "Standard", Price: 20, StorageOverageRate: 0.125, BandwidthOverageRate: 0.08, ReplicaRate: 1, PitrRate: 0.2, Cpu: 2000, Memory: 2048, MaxConnections: 200, IncludedStorage: 25, IncludedBandwidth: 200, Enabled: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseSpecification
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Slug != model.Slug {
		t.Errorf("Expected Slug %v, got %v", model.Slug, result.Slug)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Price != model.Price {
		t.Errorf("Expected Price %v, got %v", model.Price, result.Price)
	}
	if result.StorageOverageRate != model.StorageOverageRate {
		t.Errorf("Expected StorageOverageRate %v, got %v", model.StorageOverageRate, result.StorageOverageRate)
	}
	if result.BandwidthOverageRate != model.BandwidthOverageRate {
		t.Errorf("Expected BandwidthOverageRate %v, got %v", model.BandwidthOverageRate, result.BandwidthOverageRate)
	}
	if result.ReplicaRate != model.ReplicaRate {
		t.Errorf("Expected ReplicaRate %v, got %v", model.ReplicaRate, result.ReplicaRate)
	}
	if result.PitrRate != model.PitrRate {
		t.Errorf("Expected PitrRate %v, got %v", model.PitrRate, result.PitrRate)
	}
	if result.Cpu != model.Cpu {
		t.Errorf("Expected Cpu %v, got %v", model.Cpu, result.Cpu)
	}
	if result.Memory != model.Memory {
		t.Errorf("Expected Memory %v, got %v", model.Memory, result.Memory)
	}
	if result.MaxConnections != model.MaxConnections {
		t.Errorf("Expected MaxConnections %v, got %v", model.MaxConnections, result.MaxConnections)
	}
	if result.IncludedStorage != model.IncludedStorage {
		t.Errorf("Expected IncludedStorage %v, got %v", model.IncludedStorage, result.IncludedStorage)
	}
	if result.IncludedBandwidth != model.IncludedBandwidth {
		t.Errorf("Expected IncludedBandwidth %v, got %v", model.IncludedBandwidth, result.IncludedBandwidth)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
}
