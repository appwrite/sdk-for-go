package models

import (
    "encoding/json"
    "testing"
)

func TestBackupArchiveModel(t *testing.T) {
    model := BackupArchive{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        PolicyId: "did8jx6ws45jana098ab7",        Size: 100000,        Status: "completed",        StartedAt: "2020-10-15T06:38:00.000+00:00",        MigrationId: "did8jx6ws45jana098ab7",        Services: []string{"test"},        Resources: []string{"test"},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BackupArchive
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
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.PolicyId != model.PolicyId {
        t.Errorf("Expected PolicyId %v, got %v", model.PolicyId, result.PolicyId)
    }
    if result.Size != model.Size {
        t.Errorf("Expected Size %v, got %v", model.Size, result.Size)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.StartedAt != model.StartedAt {
        t.Errorf("Expected StartedAt %v, got %v", model.StartedAt, result.StartedAt)
    }
    if result.MigrationId != model.MigrationId {
        t.Errorf("Expected MigrationId %v, got %v", model.MigrationId, result.MigrationId)
    }}
