package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseBackupModel(t *testing.T) {
	model := DedicatedDatabaseBackup{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", DatabaseId: "5e5ea5c16897e", ProjectId: "5e5ea5c16897e", PolicyId: "5e5ea5c16897e", Trigger: "schedule", Type: "full", RequestedType: "incremental", FallbackReason: "PostgreSQL incremental backups are not offered because they cannot be restored: archived WAL is physical and cannot replay onto a logically-restored base. A full backup was taken instead; use a point-in-time restore (targetTime) to recover to a moment between fulls.", Status: "completed", SizeBytes: 1073741824, Error: "string"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseBackup
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
	if result.PolicyId != model.PolicyId {
		t.Errorf("Expected PolicyId %v, got %v", model.PolicyId, result.PolicyId)
	}
	if result.Trigger != model.Trigger {
		t.Errorf("Expected Trigger %v, got %v", model.Trigger, result.Trigger)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
	if result.RequestedType != model.RequestedType {
		t.Errorf("Expected RequestedType %v, got %v", model.RequestedType, result.RequestedType)
	}
	if result.FallbackReason != model.FallbackReason {
		t.Errorf("Expected FallbackReason %v, got %v", model.FallbackReason, result.FallbackReason)
	}
	if result.Status != model.Status {
		t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
	}
	if result.SizeBytes != model.SizeBytes {
		t.Errorf("Expected SizeBytes %v, got %v", model.SizeBytes, result.SizeBytes)
	}
	if result.Error != model.Error {
		t.Errorf("Expected Error %v, got %v", model.Error, result.Error)
	}
}
