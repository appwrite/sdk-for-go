package models

import (
    "encoding/json"
    "errors"
)

// Backup Model
type BackupPolicy struct {
    // Backup policy ID.
    Id string `json:"$id"`
    // Backup policy name.
    Name string `json:"name"`
    // Policy creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Policy update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // The services that are backed up by this policy.
    Services []string `json:"services"`
    // The resources that are backed up by this policy.
    Resources []string `json:"resources"`
    // The resource ID to backup. Set only if this policy should backup a single
    // resource.
    ResourceId string `json:"resourceId"`
    // The resource type to backup. Set only if this policy should backup a single
    // resource.
    ResourceType string `json:"resourceType"`
    // How many days to keep the backup before it will be automatically deleted.
    Retention int `json:"retention"`
    // Policy backup schedule in CRON format.
    Schedule string `json:"schedule"`
    // Is this policy enabled.
    Enabled bool `json:"enabled"`

    // Used by Decode() method
    data []byte
}

func (model BackupPolicy) New(data []byte) *BackupPolicy {
    model.data = data
    return &model
}

func (model *BackupPolicy) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}