package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDatabaseStatusReplicaModel(t *testing.T) {
	model := DatabaseStatusReplica{Index: 0, Role: "primary", Healthy: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DatabaseStatusReplica
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Index != model.Index {
		t.Errorf("Expected Index %v, got %v", model.Index, result.Index)
	}
	if result.Role != model.Role {
		t.Errorf("Expected Role %v, got %v", model.Role, result.Role)
	}
	if result.Healthy != model.Healthy {
		t.Errorf("Expected Healthy %v, got %v", model.Healthy, result.Healthy)
	}
}

func TestDatabaseStatusReplicaModelDefaults(t *testing.T) {
	var fresh DatabaseStatusReplica
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Index)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing index to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"unknown\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Role)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing role to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Healthy)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing healthy to default to %v, got %v", want, got)
		}
	}

	reused := DatabaseStatusReplica{Index: 0, Role: "primary", Healthy: true}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
