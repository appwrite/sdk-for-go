package models

import (
    "encoding/json"
    "testing"
)

func TestBackupRestorationModel(t *testing.T) {
    model := BackupRestoration{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        ArchiveId: "did8jx6ws45jana098ab7",        PolicyId: "did8jx6ws45jana098ab7",        Status: "completed",        StartedAt: "2020-10-15T06:38:00.000+00:00",        MigrationId: "did8jx6ws45jana098ab7",        Services: []string{"test"},        Resources: []string{"test"},        Options: "{databases.database[{oldId, newId, newName}]}",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BackupRestoration
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
    if result.ArchiveId != model.ArchiveId {
        t.Errorf("Expected ArchiveId %v, got %v", model.ArchiveId, result.ArchiveId)
    }
    if result.PolicyId != model.PolicyId {
        t.Errorf("Expected PolicyId %v, got %v", model.PolicyId, result.PolicyId)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.StartedAt != model.StartedAt {
        t.Errorf("Expected StartedAt %v, got %v", model.StartedAt, result.StartedAt)
    }
    if result.MigrationId != model.MigrationId {
        t.Errorf("Expected MigrationId %v, got %v", model.MigrationId, result.MigrationId)
    }
    if result.Options != model.Options {
        t.Errorf("Expected Options %v, got %v", model.Options, result.Options)
    }}
