package models

import (
    "encoding/json"
    "testing"
)

func TestBackupArchiveListModel(t *testing.T) {
    model := BackupArchiveList{        Total: 5,        Archives: []BackupArchive{BackupArchive{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        PolicyId: "did8jx6ws45jana098ab7",        Size: 100000,        Status: "completed",        StartedAt: "2020-10-15T06:38:00.000+00:00",        MigrationId: "did8jx6ws45jana098ab7",        Services: []string{"test"},        Resources: []string{"test"},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BackupArchiveList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
