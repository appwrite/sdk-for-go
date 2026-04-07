package models

import (
    "encoding/json"
    "testing"
)

func TestBackupRestorationListModel(t *testing.T) {
    model := BackupRestorationList{        Total: 5,        Restorations: []BackupRestoration{BackupRestoration{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        ArchiveId: "did8jx6ws45jana098ab7",        PolicyId: "did8jx6ws45jana098ab7",        Status: "completed",        StartedAt: "2020-10-15T06:38:00.000+00:00",        MigrationId: "did8jx6ws45jana098ab7",        Services: []string{"test"},        Resources: []string{"test"},        Options: "{databases.database[{oldId, newId, newName}]}",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BackupRestorationList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
