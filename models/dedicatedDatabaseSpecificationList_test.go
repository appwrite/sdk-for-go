package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseSpecificationListModel(t *testing.T) {
	model := DedicatedDatabaseSpecificationList{Specifications: []DedicatedDatabaseSpecification{DedicatedDatabaseSpecification{Slug: "s-2vcpu-2gb", Name: "Standard", Price: 20, StorageOverageRate: 0.125, BandwidthOverageRate: 0.08, ReplicaRate: 1, PitrRate: 0.2, Cpu: 2000, Memory: 2048, MaxConnections: 200, IncludedStorage: 25, IncludedBandwidth: 200, Enabled: true}}, Total: 9}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseSpecificationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestDedicatedDatabaseSpecificationListModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseSpecificationList
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

	reused := DedicatedDatabaseSpecificationList{Specifications: []DedicatedDatabaseSpecification{DedicatedDatabaseSpecification{Slug: "s-2vcpu-2gb", Name: "Standard", Price: 20, StorageOverageRate: 0.125, BandwidthOverageRate: 0.08, ReplicaRate: 1, PitrRate: 0.2, Cpu: 2000, Memory: 2048, MaxConnections: 200, IncludedStorage: 25, IncludedBandwidth: 200, Enabled: true}}, Total: 9}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
