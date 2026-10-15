package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseReplicasModel(t *testing.T) {
	model := DedicatedDatabaseReplicas{Replicas: 2, SyncMode: "async", SyncDegraded: true, SyncAcknowledgements: 1, SyncStandbyCount: 2, Members: []DedicatedDatabaseMember{DedicatedDatabaseMember{Id: "1", Role: "replica", Status: "active"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseReplicas
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Replicas != model.Replicas {
		t.Errorf("Expected Replicas %v, got %v", model.Replicas, result.Replicas)
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

func TestDedicatedDatabaseReplicasModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseReplicas
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Replicas)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing replicas to default to %v, got %v", want, got)
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

	reused := DedicatedDatabaseReplicas{Replicas: 2, SyncMode: "async", SyncDegraded: true, SyncAcknowledgements: 1, SyncStandbyCount: 2, Members: []DedicatedDatabaseMember{DedicatedDatabaseMember{Id: "1", Role: "replica", Status: "active"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
