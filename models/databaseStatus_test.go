package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDatabaseStatusModel(t *testing.T) {
	model := DatabaseStatus{Health: "healthy", Ready: true, Engine: "postgresql", Version: "17", Uptime: 86400, SyncMode: "async", SyncDegraded: true, SyncAcknowledgements: 1, SyncStandbyCount: 2, Replicas: []DatabaseStatusReplica{DatabaseStatusReplica{Index: 0, Role: "primary", Healthy: true}}, Volumes: []DatabaseStatusVolume{DatabaseStatusVolume{Path: "/var/lib/postgresql/data", UsedPercent: "45%", Available: "55GB", Mounted: true}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DatabaseStatus
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Health != model.Health {
		t.Errorf("Expected Health %v, got %v", model.Health, result.Health)
	}
	if result.Ready != model.Ready {
		t.Errorf("Expected Ready %v, got %v", model.Ready, result.Ready)
	}
	if result.Engine != model.Engine {
		t.Errorf("Expected Engine %v, got %v", model.Engine, result.Engine)
	}
	if result.Version != model.Version {
		t.Errorf("Expected Version %v, got %v", model.Version, result.Version)
	}
	if result.Uptime != model.Uptime {
		t.Errorf("Expected Uptime %v, got %v", model.Uptime, result.Uptime)
	}
	if result.SyncMode != model.SyncMode {
		t.Errorf("Expected SyncMode %v, got %v", model.SyncMode, result.SyncMode)
	}
	if result.SyncDegraded != model.SyncDegraded {
		t.Errorf("Expected SyncDegraded %v, got %v", model.SyncDegraded, result.SyncDegraded)
	}
	if result.SyncAcknowledgements != model.SyncAcknowledgements {
		t.Errorf("Expected SyncAcknowledgements %v, got %v", model.SyncAcknowledgements, result.SyncAcknowledgements)
	}
	if result.SyncStandbyCount != model.SyncStandbyCount {
		t.Errorf("Expected SyncStandbyCount %v, got %v", model.SyncStandbyCount, result.SyncStandbyCount)
	}
}

func TestDatabaseStatusModelDefaults(t *testing.T) {
	var fresh DatabaseStatus
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"unknown\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Health)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing health to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Ready)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ready to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"postgresql\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Engine)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing engine to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Version)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing version to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Uptime)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing uptime to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"async\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SyncMode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing syncMode to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SyncDegraded)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing syncDegraded to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SyncAcknowledgements)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing syncAcknowledgements to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SyncStandbyCount)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing syncStandbyCount to default to %v, got %v", want, got)
		}
	}

	reused := DatabaseStatus{Health: "healthy", Ready: true, Engine: "postgresql", Version: "17", Uptime: 86400, SyncMode: "async", SyncDegraded: true, SyncAcknowledgements: 1, SyncStandbyCount: 2, Replicas: []DatabaseStatusReplica{DatabaseStatusReplica{Index: 0, Role: "primary", Healthy: true}}, Volumes: []DatabaseStatusVolume{DatabaseStatusVolume{Path: "/var/lib/postgresql/data", UsedPercent: "45%", Available: "55GB", Mounted: true}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
