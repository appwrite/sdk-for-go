package models

import (
	"encoding/json"
	"reflect"
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

func TestDedicatedDatabaseSpecificationModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseSpecification
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Slug)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing slug to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Name)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing name to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Price)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing price to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.StorageOverageRate)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing storageOverageRate to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BandwidthOverageRate)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing bandwidthOverageRate to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ReplicaRate)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing replicaRate to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PitrRate)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing pitrRate to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Cpu)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing cpu to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Memory)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing memory to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.MaxConnections)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing maxConnections to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.IncludedStorage)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing includedStorage to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.IncludedBandwidth)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing includedBandwidth to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Enabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing enabled to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabaseSpecification{Slug: "s-2vcpu-2gb", Name: "Standard", Price: 20, StorageOverageRate: 0.125, BandwidthOverageRate: 0.08, ReplicaRate: 1, PitrRate: 0.2, Cpu: 2000, Memory: 2048, MaxConnections: 200, IncludedStorage: 25, IncludedBandwidth: 200, Enabled: true}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
