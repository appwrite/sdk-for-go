package models

import (
    "encoding/json"
    "testing"
)

func TestBackupPolicyModel(t *testing.T) {
    model := BackupPolicy{        Id: "5e5ea5c16897e",        Name: "Hourly backups",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Services: []string{"test"},        Resources: []string{"test"},        Retention: 7,        Schedule: "0 * * * *",        Enabled: true,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BackupPolicy
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.Retention != model.Retention {
        t.Errorf("Expected Retention %v, got %v", model.Retention, result.Retention)
    }
    if result.Schedule != model.Schedule {
        t.Errorf("Expected Schedule %v, got %v", model.Schedule, result.Schedule)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }}
