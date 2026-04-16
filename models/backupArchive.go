package models

import (
    "encoding/json"
    "errors"
)

// Archive Model
type BackupArchive struct {
    // Archive ID.
    Id string `json:"$id"`
    // Archive creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Archive update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Archive policy ID.
    PolicyId string `json:"policyId"`
    // Archive size in bytes.
    Size int `json:"size"`
    // The status of the archive creation. Possible values: pending, processing,
    // uploading, completed, failed.
    Status string `json:"status"`
    // The backup start time.
    StartedAt string `json:"startedAt"`
    // Migration ID.
    MigrationId string `json:"migrationId"`
    // The services that are backed up by this archive.
    Services []string `json:"services"`
    // The resources that are backed up by this archive.
    Resources []string `json:"resources"`
    // The resource ID to backup. Set only if this archive should backup a single
    // resource.
    ResourceId string `json:"resourceId"`
    // The resource type to backup. Set only if this archive should backup a
    // single resource.
    ResourceType string `json:"resourceType"`

    // Used by Decode() method
    data []byte
}

func (model BackupArchive) New(data []byte) *BackupArchive {
    model.data = data
    return &model
}

func (model *BackupArchive) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}