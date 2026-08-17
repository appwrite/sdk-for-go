package models

import (
	"encoding/json"
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
