package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabasePoolerModel(t *testing.T) {
	model := DedicatedDatabasePooler{Enabled: true, Mode: "transaction", MaxConnections: 200, DefaultPoolSize: 25, Port: 6432, ReadWriteSplitting: true, PoolerCpuRequest: "100m", PoolerCpuLimit: "200m", PoolerMemoryRequest: "64Mi", PoolerMemoryLimit: "128Mi"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabasePooler
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
	if result.Mode != model.Mode {
		t.Errorf("Expected Mode %v, got %v", model.Mode, result.Mode)
	}
	if result.MaxConnections != model.MaxConnections {
		t.Errorf("Expected MaxConnections %v, got %v", model.MaxConnections, result.MaxConnections)
	}
	if result.DefaultPoolSize != model.DefaultPoolSize {
		t.Errorf("Expected DefaultPoolSize %v, got %v", model.DefaultPoolSize, result.DefaultPoolSize)
	}
	if result.Port != model.Port {
		t.Errorf("Expected Port %v, got %v", model.Port, result.Port)
	}
	if result.ReadWriteSplitting != model.ReadWriteSplitting {
		t.Errorf("Expected ReadWriteSplitting %v, got %v", model.ReadWriteSplitting, result.ReadWriteSplitting)
	}
	if result.PoolerCpuRequest != model.PoolerCpuRequest {
		t.Errorf("Expected PoolerCpuRequest %v, got %v", model.PoolerCpuRequest, result.PoolerCpuRequest)
	}
	if result.PoolerCpuLimit != model.PoolerCpuLimit {
		t.Errorf("Expected PoolerCpuLimit %v, got %v", model.PoolerCpuLimit, result.PoolerCpuLimit)
	}
	if result.PoolerMemoryRequest != model.PoolerMemoryRequest {
		t.Errorf("Expected PoolerMemoryRequest %v, got %v", model.PoolerMemoryRequest, result.PoolerMemoryRequest)
	}
	if result.PoolerMemoryLimit != model.PoolerMemoryLimit {
		t.Errorf("Expected PoolerMemoryLimit %v, got %v", model.PoolerMemoryLimit, result.PoolerMemoryLimit)
	}
}

func TestDedicatedDatabasePoolerModelDefaults(t *testing.T) {
	var fresh DedicatedDatabasePooler
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
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
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"transaction\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Mode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing mode to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("100"), &want); err != nil {
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
		if err := json.Unmarshal([]byte("20"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DefaultPoolSize)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing defaultPoolSize to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("6432"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Port)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing port to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ReadWriteSplitting)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing readWriteSplitting to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PoolerCpuRequest)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing poolerCpuRequest to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PoolerCpuLimit)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing poolerCpuLimit to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PoolerMemoryRequest)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing poolerMemoryRequest to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.PoolerMemoryLimit)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing poolerMemoryLimit to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabasePooler{Enabled: true, Mode: "transaction", MaxConnections: 200, DefaultPoolSize: 25, Port: 6432, ReadWriteSplitting: true, PoolerCpuRequest: "100m", PoolerCpuLimit: "200m", PoolerMemoryRequest: "64Mi", PoolerMemoryLimit: "128Mi"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
