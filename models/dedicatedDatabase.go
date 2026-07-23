package models

import (
    "encoding/json"
    "errors"
)

// DedicatedDatabase Model
type DedicatedDatabase struct {
    // Dedicated database ID.
    Id string `json:"$id"`
    // Database creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Database update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Project ID that owns this database.
    ProjectId string `json:"projectId"`
    // Database display name.
    Name string `json:"name"`
    // Product API that owns this database: tablesdb, documentsdb, vectorsdb,
    // mysql, postgresql, or mongodb.
    Api string `json:"api"`
    // Database engine: postgresql, mysql, mariadb, or mongodb.
    Engine string `json:"engine"`
    // Database engine version.
    Version string `json:"version"`
    // Specification identifier.
    Specification string `json:"specification"`
    // Database backend provider. Possible values: prisma, edge.
    Backend string `json:"backend"`
    // Database hostname for connections.
    Hostname string `json:"hostname"`
    // Database port for connections.
    ConnectionPort int `json:"connectionPort"`
    // Database username for connections.
    ConnectionUser string `json:"connectionUser"`
    // Database password for connections.
    ConnectionPassword string `json:"connectionPassword"`
    // Full database connection string (URI format).
    ConnectionString string `json:"connectionString"`
    // Whether SSL/TLS is required for client connections.
    Ssl bool `json:"ssl"`
    // Database status. Possible values: provisioning, ready, inactive, paused,
    // failed, deleted, restoring, scaling.
    Status string `json:"status"`
    // Container status for lifecycle-managed database runtimes: active or
    // inactive.
    ContainerStatus string `json:"containerStatus"`
    // Last activity timestamp in ISO 8601 format.
    LastAccessedAt string `json:"lastAccessedAt"`
    // Display-only timestamp when the database is expected to be considered idle
    // (ISO 8601 format). Derived from last activity; lifecycle transitions are
    // driven by lifecycleState.
    IdleUntil string `json:"idleUntil"`
    // Idle-lifecycle state of the database. Possible values: active, warm, cold,
    // hibernated.
    LifecycleState string `json:"lifecycleState"`
    // Minutes of inactivity before container scales to zero.
    IdleTimeoutMinutes int `json:"idleTimeoutMinutes"`
    // CPU allocated in millicores.
    Cpu int `json:"cpu"`
    // Memory allocated in MB.
    Memory int `json:"memory"`
    // Storage allocated in GB.
    Storage int `json:"storage"`
    // Storage class. Currently always 'ssd'; DigitalOcean exposes a single
    // block-storage class.
    StorageClass string `json:"storageClass"`
    // Maximum storage allowed in GB. 0 means use system default.
    StorageMaxGb int `json:"storageMaxGb"`
    // Kubernetes node pool where the database is scheduled.
    NodePool string `json:"nodePool"`
    // Number of high availability replicas. High availability is enabled when
    // greater than 0.
    Replicas int `json:"replicas"`
    // Replication sync mode: async, sync, or quorum.
    SyncMode string `json:"syncMode"`
    // Number of cross-region replicas. Cross-region availability is enabled when
    // greater than 0.
    CrossRegionReplicas int `json:"crossRegionReplicas"`
    // Maximum concurrent connections.
    NetworkMaxConnections int `json:"networkMaxConnections"`
    // Connection idle timeout in seconds.
    NetworkIdleTimeoutSeconds int `json:"networkIdleTimeoutSeconds"`
    // IP addresses/CIDR ranges allowed to connect.
    NetworkIPAllowlist []string `json:"networkIPAllowlist"`
    // Whether automatic backups are enabled.
    BackupEnabled bool `json:"backupEnabled"`
    // Whether point-in-time recovery is enabled.
    Pitr bool `json:"pitr"`
    // Number of days to retain PITR data.
    PitrRetentionDays int `json:"pitrRetentionDays"`
    // Whether automatic storage expansion is enabled.
    StorageAutoscaling bool `json:"storageAutoscaling"`
    // Storage usage percentage that triggers automatic expansion.
    StorageAutoscalingThresholdPercent int `json:"storageAutoscalingThresholdPercent"`
    // Maximum storage size in GB for autoscaling. 0 means no limit.
    StorageAutoscalingMaxGb int `json:"storageAutoscalingMaxGb"`
    // Day of the week for the maintenance window. Possible values: sun, mon, tue,
    // wed, thu, fri, sat.
    MaintenanceWindowDay string `json:"maintenanceWindowDay"`
    // Hour in UTC (0-23) when the maintenance window starts.
    MaintenanceWindowHourUtc int `json:"maintenanceWindowHourUtc"`
    // Whether metrics collection is enabled.
    MetricsEnabled bool `json:"metricsEnabled"`
    // Whether the SQL API sidecar is enabled for this database.
    SqlApiEnabled bool `json:"sqlApiEnabled"`
    // Statement types accepted by the SQL API. Defaults to read/write DML only;
    // DDL/DCL types (CREATE, ALTER, DROP, TRUNCATE, GRANT, REVOKE) are opt-in per
    // database. Allowed values: SELECT, INSERT, UPDATE, DELETE, CREATE, ALTER,
    // DROP, TRUNCATE, GRANT, REVOKE.
    SqlApiAllowedStatements []string `json:"sqlApiAllowedStatements"`
    // Maximum rows returned per SQL API execution. Results larger than this are
    // truncated.
    SqlApiMaxRows int `json:"sqlApiMaxRows"`
    // Maximum serialised SQL API result payload in bytes. Results larger than
    // this are truncated.
    SqlApiMaxBytes int `json:"sqlApiMaxBytes"`
    // Maximum server-side SQL API execution time in seconds before the query is
    // cancelled.
    SqlApiTimeoutSeconds int `json:"sqlApiTimeoutSeconds"`
    // Error message if status is failed.
    Error string `json:"error"`

    // Used by Decode() method
    data []byte
}

func (model DedicatedDatabase) New(data []byte) *DedicatedDatabase {
    model.data = data
    return &model
}

func (model *DedicatedDatabase) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}