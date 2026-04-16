package models

import (
    "encoding/json"
    "testing"
)

func TestBackupPolicyListModel(t *testing.T) {
    model := BackupPolicyList{        Total: 5,        Policies: []BackupPolicy{BackupPolicy{        Id: "5e5ea5c16897e",        Name: "Hourly backups",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Services: []string{"test"},        Resources: []string{"test"},        Retention: 7,        Schedule: "0 * * * *",        Enabled: true,    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result BackupPolicyList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
