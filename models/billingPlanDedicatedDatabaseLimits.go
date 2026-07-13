package models

import (
    "encoding/json"
    "errors"
)

// DedicatedDatabaseLimits Model
type BillingPlanDedicatedDatabaseLimits struct {
    // Minimum CPU allocation in millicores.
    MinCpu int `json:"minCpu"`
    // Maximum CPU allocation in millicores.
    MaxCpu int `json:"maxCpu"`
    // Minimum memory allocation in megabytes.
    MinMemoryMb int `json:"minMemoryMb"`
    // Maximum memory allocation in megabytes.
    MaxMemoryMb int `json:"maxMemoryMb"`
    // Minimum storage allocation in gigabytes.
    MinStorageGb int `json:"minStorageGb"`
    // Maximum storage allocation in gigabytes.
    MaxStorageGb int `json:"maxStorageGb"`
    // Maximum number of high-availability replicas per dedicated database.
    MaxReplicas int `json:"maxReplicas"`
    // Maximum number of client connections.
    MaxConnections int `json:"maxConnections"`
    // Maximum number of entries allowed in the IP allowlist.
    MaxIpAllowlistSize int `json:"maxIpAllowlistSize"`
    // Maximum number of database extensions that can be enabled.
    MaxExtensions int `json:"maxExtensions"`
    // Maximum number of days a backup can be retained.
    MaxBackupRetentionDays int `json:"maxBackupRetentionDays"`
    // Maximum number of days of point-in-time recovery data that can be retained.
    MaxPitrRetentionDays int `json:"maxPitrRetentionDays"`
    // Maximum number of rows a single SQL API query can return.
    MaxSqlApiMaxRows int `json:"maxSqlApiMaxRows"`
    // Maximum response size in bytes for a single SQL API query.
    MaxSqlApiMaxBytes int `json:"maxSqlApiMaxBytes"`
    // Maximum execution time in seconds for a single SQL API query.
    MaxSqlApiTimeoutSeconds int `json:"maxSqlApiTimeoutSeconds"`
    // Maximum number of SQL statement types that can be permitted through the SQL
    // API.
    MaxSqlApiAllowedStatements int `json:"maxSqlApiAllowedStatements"`
    // SQL statement types permitted through the SQL API.
    AllowedSqlStatements []string `json:"allowedSqlStatements"`
    // Storage classes available for dedicated databases.
    AllowedStorageClasses []string `json:"allowedStorageClasses"`
    // Replica synchronization modes available for dedicated databases.
    AllowedSyncModes []string `json:"allowedSyncModes"`

    // Used by Decode() method
    data []byte
}

func (model BillingPlanDedicatedDatabaseLimits) New(data []byte) *BillingPlanDedicatedDatabaseLimits {
    model.data = data
    return &model
}

func (model *BillingPlanDedicatedDatabaseLimits) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}