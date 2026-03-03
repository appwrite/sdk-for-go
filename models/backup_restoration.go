package models

import (
    "encoding/json"
    "errors"
)

// Restoration Model
type BackupRestoration struct {
    // Restoration ID.
    Id string `json:"$id"`
    // Restoration creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Restoration update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Backup archive ID.
    ArchiveId string `json:"archiveId"`
    // Backup policy ID.
    PolicyId string `json:"policyId"`
    // The status of the restoration. Possible values: pending, downloading,
    // processing, completed, failed.
    Status string `json:"status"`
    // The backup start time.
    StartedAt string `json:"startedAt"`
    // Migration ID.
    MigrationId string `json:"migrationId"`
    // The services that are backed up by this policy.
    Services []string `json:"services"`
    // The resources that are backed up by this policy.
    Resources []string `json:"resources"`
    // Optional data in key-value object.
    Options string `json:"options"`

    // Used by Decode() method
    data []byte
}

func (model BackupRestoration) New(data []byte) *BackupRestoration {
    model.data = data
    return &model
}

func (model *BackupRestoration) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}