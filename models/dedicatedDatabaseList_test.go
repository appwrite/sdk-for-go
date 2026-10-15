package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseListModel(t *testing.T) {
	model := DedicatedDatabaseList{Total: 5, Databases: []DedicatedDatabase{DedicatedDatabase{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", Name: "My Production Database", Version: "16", Specification: "s-2vcpu-2gb", Backend: "edge", Hostname: "db-myproject-mydb.fra.appwrite.center", ConnectionUser: "appwrite_user", ConnectionPassword: "••••••••", ConnectionString: "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require", Ssl: true, Status: "ready", ContainerStatus: "active", LifecycleState: "active", IdleTimeoutMinutes: 15, Cpu: 2000, Memory: 4096, Storage: 100, StorageStatus: "resizing", StorageTargetGb: 120, StorageClass: "ssd", StorageMaxGb: 100, NodePool: "db-pool-4vcpu-8gb", Replicas: 2, SyncMode: "async", NetworkMaxConnections: 500, NetworkIdleTimeoutSeconds: 900, NetworkIPAllowlist: []string{"test"}, BackupEnabled: true, Pitr: true, PitrRetentionDays: 14, StorageAutoscaling: true, StorageAutoscalingThresholdPercent: 85, StorageAutoscalingMaxGb: 30, MaintenanceWindowDay: "sun", MaintenanceWindowHourUtc: 3, MetricsEnabled: true, SqlApiEnabled: true, SqlApiAllowedStatements: []string{"test"}, SqlApiMaxRows: 10000, SqlApiMaxBytes: 10485760, SqlApiTimeoutSeconds: 30, Error: "string"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestDedicatedDatabaseListModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabaseList{Total: 5, Databases: []DedicatedDatabase{DedicatedDatabase{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", Name: "My Production Database", Version: "16", Specification: "s-2vcpu-2gb", Backend: "edge", Hostname: "db-myproject-mydb.fra.appwrite.center", ConnectionUser: "appwrite_user", ConnectionPassword: "••••••••", ConnectionString: "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require", Ssl: true, Status: "ready", ContainerStatus: "active", LifecycleState: "active", IdleTimeoutMinutes: 15, Cpu: 2000, Memory: 4096, Storage: 100, StorageStatus: "resizing", StorageTargetGb: 120, StorageClass: "ssd", StorageMaxGb: 100, NodePool: "db-pool-4vcpu-8gb", Replicas: 2, SyncMode: "async", NetworkMaxConnections: 500, NetworkIdleTimeoutSeconds: 900, NetworkIPAllowlist: []string{"test"}, BackupEnabled: true, Pitr: true, PitrRetentionDays: 14, StorageAutoscaling: true, StorageAutoscalingThresholdPercent: 85, StorageAutoscalingMaxGb: 30, MaintenanceWindowDay: "sun", MaintenanceWindowHourUtc: 3, MetricsEnabled: true, SqlApiEnabled: true, SqlApiAllowedStatements: []string{"test"}, SqlApiMaxRows: 10000, SqlApiMaxBytes: 10485760, SqlApiTimeoutSeconds: 30, Error: "string"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
