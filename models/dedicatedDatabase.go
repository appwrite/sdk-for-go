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
	Api *string `json:"api"`
	// Database engine: postgresql, mysql, or mongodb. Null until the backing
	// reports one.
	Engine *string `json:"engine"`
	// Database engine version.
	Version string `json:"version"`
	// Specification identifier.
	Specification string `json:"specification"`
	// Database backend provider. Possible values: edge.
	Backend string `json:"backend"`
	// Database hostname for connections.
	Hostname string `json:"hostname"`
	// Database port for connections. Derived from the engine when the backing has
	// not reported one yet.
	ConnectionPort *int `json:"connectionPort"`
	// Database username for connections.
	ConnectionUser string `json:"connectionUser"`
	// Database password for connections.
	ConnectionPassword string `json:"connectionPassword"`
	// Committed generation of the primary connection credentials. Null until the
	// rotation contract has been initialized.
	CredentialGeneration *int `json:"credentialGeneration"`
	// Full database connection string (URI format).
	ConnectionString string `json:"connectionString"`
	// Whether SSL/TLS is required for client connections.
	Ssl bool `json:"ssl"`
	// Database status. Possible values: provisioning, ready, inactive, paused,
	// failed, deleted, restoring, scaling.
	Status string `json:"status"`
	// Container status for lifecycle-managed database runtimes: active, inactive,
	// or failed (a wake was declined because the backing namespace is gone).
	ContainerStatus string `json:"containerStatus"`
	// Last activity timestamp in ISO 8601 format.
	LastAccessedAt *string `json:"lastAccessedAt"`
	// Display-only timestamp when the database is expected to be considered idle
	// (ISO 8601 format). Derived from last activity; lifecycle transitions are
	// driven by lifecycleState.
	IdleUntil *string `json:"idleUntil"`
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
	// Storage resize status. Possible values: idle (no resize in flight),
	// resizing (the volume is growing towards storageTargetGb).
	StorageStatus string `json:"storageStatus"`
	// Size in GB the volume is growing towards while storageStatus is resizing. 0
	// when no resize is in flight.
	StorageTargetGb int `json:"storageTargetGb"`
	// Time the in-flight storage resize started, in ISO 8601 format.
	StorageResizeStartedAt *string `json:"storageResizeStartedAt"`
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
	// Maximum concurrent client connections. This is the limit a client pool may
	// reach; the engine's own max_connections reported by the status endpoint is
	// a smaller backend limit the pooler multiplexes onto and does not constrain
	// a client pool.
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
	// Maximum storage size in GB for autoscaling. Defaults to 3 times the
	// specification's storage. 0 means no limit.
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

// UnmarshalJSON decodes a response into a fresh model: present keys keep their
// value, and required keys the response omits take their documented default.
func (model *DedicatedDatabase) UnmarshalJSON(data []byte) error {
	type alias DedicatedDatabase
	decoded := alias{data: model.data}
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}

	var present struct {
		Id                                 json.RawMessage `json:"$id"`
		CreatedAt                          json.RawMessage `json:"$createdAt"`
		UpdatedAt                          json.RawMessage `json:"$updatedAt"`
		ProjectId                          json.RawMessage `json:"projectId"`
		Name                               json.RawMessage `json:"name"`
		Version                            json.RawMessage `json:"version"`
		Specification                      json.RawMessage `json:"specification"`
		Backend                            json.RawMessage `json:"backend"`
		Hostname                           json.RawMessage `json:"hostname"`
		ConnectionUser                     json.RawMessage `json:"connectionUser"`
		ConnectionPassword                 json.RawMessage `json:"connectionPassword"`
		ConnectionString                   json.RawMessage `json:"connectionString"`
		Ssl                                json.RawMessage `json:"ssl"`
		Status                             json.RawMessage `json:"status"`
		ContainerStatus                    json.RawMessage `json:"containerStatus"`
		LifecycleState                     json.RawMessage `json:"lifecycleState"`
		IdleTimeoutMinutes                 json.RawMessage `json:"idleTimeoutMinutes"`
		Cpu                                json.RawMessage `json:"cpu"`
		Memory                             json.RawMessage `json:"memory"`
		Storage                            json.RawMessage `json:"storage"`
		StorageStatus                      json.RawMessage `json:"storageStatus"`
		StorageTargetGb                    json.RawMessage `json:"storageTargetGb"`
		StorageClass                       json.RawMessage `json:"storageClass"`
		StorageMaxGb                       json.RawMessage `json:"storageMaxGb"`
		NodePool                           json.RawMessage `json:"nodePool"`
		Replicas                           json.RawMessage `json:"replicas"`
		SyncMode                           json.RawMessage `json:"syncMode"`
		NetworkMaxConnections              json.RawMessage `json:"networkMaxConnections"`
		NetworkIdleTimeoutSeconds          json.RawMessage `json:"networkIdleTimeoutSeconds"`
		NetworkIPAllowlist                 json.RawMessage `json:"networkIPAllowlist"`
		BackupEnabled                      json.RawMessage `json:"backupEnabled"`
		Pitr                               json.RawMessage `json:"pitr"`
		PitrRetentionDays                  json.RawMessage `json:"pitrRetentionDays"`
		StorageAutoscaling                 json.RawMessage `json:"storageAutoscaling"`
		StorageAutoscalingThresholdPercent json.RawMessage `json:"storageAutoscalingThresholdPercent"`
		StorageAutoscalingMaxGb            json.RawMessage `json:"storageAutoscalingMaxGb"`
		MaintenanceWindowDay               json.RawMessage `json:"maintenanceWindowDay"`
		MaintenanceWindowHourUtc           json.RawMessage `json:"maintenanceWindowHourUtc"`
		MetricsEnabled                     json.RawMessage `json:"metricsEnabled"`
		SqlApiEnabled                      json.RawMessage `json:"sqlApiEnabled"`
		SqlApiAllowedStatements            json.RawMessage `json:"sqlApiAllowedStatements"`
		SqlApiMaxRows                      json.RawMessage `json:"sqlApiMaxRows"`
		SqlApiMaxBytes                     json.RawMessage `json:"sqlApiMaxBytes"`
		SqlApiTimeoutSeconds               json.RawMessage `json:"sqlApiTimeoutSeconds"`
		Error                              json.RawMessage `json:"error"`
	}
	if err := json.Unmarshal(data, &present); err != nil {
		return err
	}
	if present.Id == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Id); err != nil {
			return err
		}
	}
	if present.CreatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.CreatedAt); err != nil {
			return err
		}
	}
	if present.UpdatedAt == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.UpdatedAt); err != nil {
			return err
		}
	}
	if present.ProjectId == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ProjectId); err != nil {
			return err
		}
	}
	if present.Name == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Name); err != nil {
			return err
		}
	}
	if present.Version == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Version); err != nil {
			return err
		}
	}
	if present.Specification == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Specification); err != nil {
			return err
		}
	}
	if present.Backend == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Backend); err != nil {
			return err
		}
	}
	if present.Hostname == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Hostname); err != nil {
			return err
		}
	}
	if present.ConnectionUser == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConnectionUser); err != nil {
			return err
		}
	}
	if present.ConnectionPassword == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConnectionPassword); err != nil {
			return err
		}
	}
	if present.ConnectionString == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ConnectionString); err != nil {
			return err
		}
	}
	if present.Ssl == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.Ssl); err != nil {
			return err
		}
	}
	if present.Status == nil {
		if err := json.Unmarshal([]byte("\"provisioning\""), &decoded.Status); err != nil {
			return err
		}
	}
	if present.ContainerStatus == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.ContainerStatus); err != nil {
			return err
		}
	}
	if present.LifecycleState == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.LifecycleState); err != nil {
			return err
		}
	}
	if present.IdleTimeoutMinutes == nil {
		if err := json.Unmarshal([]byte("15"), &decoded.IdleTimeoutMinutes); err != nil {
			return err
		}
	}
	if present.Cpu == nil {
		if err := json.Unmarshal([]byte("500"), &decoded.Cpu); err != nil {
			return err
		}
	}
	if present.Memory == nil {
		if err := json.Unmarshal([]byte("512"), &decoded.Memory); err != nil {
			return err
		}
	}
	if present.Storage == nil {
		if err := json.Unmarshal([]byte("10"), &decoded.Storage); err != nil {
			return err
		}
	}
	if present.StorageStatus == nil {
		if err := json.Unmarshal([]byte("\"idle\""), &decoded.StorageStatus); err != nil {
			return err
		}
	}
	if present.StorageTargetGb == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.StorageTargetGb); err != nil {
			return err
		}
	}
	if present.StorageClass == nil {
		if err := json.Unmarshal([]byte("\"ssd\""), &decoded.StorageClass); err != nil {
			return err
		}
	}
	if present.StorageMaxGb == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.StorageMaxGb); err != nil {
			return err
		}
	}
	if present.NodePool == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.NodePool); err != nil {
			return err
		}
	}
	if present.Replicas == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.Replicas); err != nil {
			return err
		}
	}
	if present.SyncMode == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.SyncMode); err != nil {
			return err
		}
	}
	if present.NetworkMaxConnections == nil {
		if err := json.Unmarshal([]byte("100"), &decoded.NetworkMaxConnections); err != nil {
			return err
		}
	}
	if present.NetworkIdleTimeoutSeconds == nil {
		if err := json.Unmarshal([]byte("900"), &decoded.NetworkIdleTimeoutSeconds); err != nil {
			return err
		}
	}
	if present.NetworkIPAllowlist == nil {
		if err := json.Unmarshal([]byte("[]"), &decoded.NetworkIPAllowlist); err != nil {
			return err
		}
	}
	if present.BackupEnabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.BackupEnabled); err != nil {
			return err
		}
	}
	if present.Pitr == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.Pitr); err != nil {
			return err
		}
	}
	if present.PitrRetentionDays == nil {
		if err := json.Unmarshal([]byte("7"), &decoded.PitrRetentionDays); err != nil {
			return err
		}
	}
	if present.StorageAutoscaling == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.StorageAutoscaling); err != nil {
			return err
		}
	}
	if present.StorageAutoscalingThresholdPercent == nil {
		if err := json.Unmarshal([]byte("85"), &decoded.StorageAutoscalingThresholdPercent); err != nil {
			return err
		}
	}
	if present.StorageAutoscalingMaxGb == nil {
		if err := json.Unmarshal([]byte("0"), &decoded.StorageAutoscalingMaxGb); err != nil {
			return err
		}
	}
	if present.MaintenanceWindowDay == nil {
		if err := json.Unmarshal([]byte("\"sun\""), &decoded.MaintenanceWindowDay); err != nil {
			return err
		}
	}
	if present.MaintenanceWindowHourUtc == nil {
		if err := json.Unmarshal([]byte("3"), &decoded.MaintenanceWindowHourUtc); err != nil {
			return err
		}
	}
	if present.MetricsEnabled == nil {
		if err := json.Unmarshal([]byte("false"), &decoded.MetricsEnabled); err != nil {
			return err
		}
	}
	if present.SqlApiEnabled == nil {
		if err := json.Unmarshal([]byte("true"), &decoded.SqlApiEnabled); err != nil {
			return err
		}
	}
	if present.SqlApiAllowedStatements == nil {
		if err := json.Unmarshal([]byte("[\"SELECT\",\"INSERT\",\"UPDATE\",\"DELETE\"]"), &decoded.SqlApiAllowedStatements); err != nil {
			return err
		}
	}
	if present.SqlApiMaxRows == nil {
		if err := json.Unmarshal([]byte("10000"), &decoded.SqlApiMaxRows); err != nil {
			return err
		}
	}
	if present.SqlApiMaxBytes == nil {
		if err := json.Unmarshal([]byte("10485760"), &decoded.SqlApiMaxBytes); err != nil {
			return err
		}
	}
	if present.SqlApiTimeoutSeconds == nil {
		if err := json.Unmarshal([]byte("30"), &decoded.SqlApiTimeoutSeconds); err != nil {
			return err
		}
	}
	if present.Error == nil {
		if err := json.Unmarshal([]byte("\"\""), &decoded.Error); err != nil {
			return err
		}
	}

	*model = DedicatedDatabase(decoded)
	return nil
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
