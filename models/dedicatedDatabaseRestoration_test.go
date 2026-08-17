package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseRestorationModel(t *testing.T) {
	model := DedicatedDatabaseRestoration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", DatabaseId: "5e5ea5c16897e", SourceDatabaseId: "5e5ea5c16897e", ProjectId: "5e5ea5c16897e", BackupId: "5e5ea5c16897e", Type: "backup", Status: "completed", TargetTime: "2020-10-15T06:38:00.000+00:00", StartedAt: "2020-10-15T06:38:00.000+00:00", CompletedAt: "2020-10-15T06:38:00.000+00:00", Error: "string"}

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
	if result.SourceDatabaseId != model.SourceDatabaseId {
		t.Errorf("Expected SourceDatabaseId %v, got %v", model.SourceDatabaseId, result.SourceDatabaseId)
	}
	if result.ProjectId != model.ProjectId {
		t.Errorf("Expected ProjectId %v, got %v", model.ProjectId, result.ProjectId)
	}
	if result.BackupId != model.BackupId {
		t.Errorf("Expected BackupId %v, got %v", model.BackupId, result.BackupId)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
	if result.Status != model.Status {
		t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
	}
	if result.TargetTime != model.TargetTime {
		t.Errorf("Expected TargetTime %v, got %v", model.TargetTime, result.TargetTime)
	}
	if result.StartedAt != model.StartedAt {
		t.Errorf("Expected StartedAt %v, got %v", model.StartedAt, result.StartedAt)
	}
	if result.CompletedAt != model.CompletedAt {
		t.Errorf("Expected CompletedAt %v, got %v", model.CompletedAt, result.CompletedAt)
	}
	if result.Error != model.Error {
		t.Errorf("Expected Error %v, got %v", model.Error, result.Error)
	}
}
