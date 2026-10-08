package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseRestorationModel(t *testing.T) {
	model := DedicatedDatabaseRestoration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", DatabaseId: "5e5ea5c16897e", ProjectId: "5e5ea5c16897e", Type: "backup", Status: "completed", Error: "string"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseRestoration
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
	if result.DatabaseId != model.DatabaseId {
		t.Errorf("Expected DatabaseId %v, got %v", model.DatabaseId, result.DatabaseId)
	}
	if result.ProjectId != model.ProjectId {
		t.Errorf("Expected ProjectId %v, got %v", model.ProjectId, result.ProjectId)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
	if result.Status != model.Status {
		t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
	}
	if result.Error != model.Error {
		t.Errorf("Expected Error %v, got %v", model.Error, result.Error)
	}
}
