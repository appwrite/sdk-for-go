package mysql

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/appwrite/sdk-for-go/v7/client"
)

func TestMysql(t *testing.T) {
	newTestClient := func(ts *httptest.Server) client.Client {
		c := client.New()
		c.Endpoint = ts.URL
		c.Client = ts.Client()
		return c
	}

	t.Run("Test List", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "databases": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "projectId": "5e5ea5c16897e",
            "name": "My Production Database",
            "api": "postgresql",
            "engine": "postgresql",
            "version": "16",
            "specification": "s-2vcpu-2gb",
            "backend": "edge",
            "hostname": "db-myproject-mydb.fra.appwrite.center",
            "connectionPort": 5432,
            "connectionUser": "appwrite_user",
            "connectionPassword": "••••••••",
            "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
            "ssl": true,
            "status": "ready",
            "containerStatus": "active",
            "lifecycleState": "active",
            "idleTimeoutMinutes": 15,
            "cpu": 2000,
            "memory": 4096,
            "storage": 100,
            "storageClass": "ssd",
            "storageMaxGb": 100,
            "nodePool": "db-pool-4vcpu-8gb",
            "replicas": 2,
            "syncMode": "async",
            "networkMaxConnections": 500,
            "networkIdleTimeoutSeconds": 900,
            "networkIPAllowlist": [],
            "backupEnabled": true,
            "pitr": true,
            "pitrRetentionDays": 14,
            "storageAutoscaling": true,
            "storageAutoscalingThresholdPercent": 85,
            "storageAutoscalingMaxGb": 500,
            "maintenanceWindowDay": "sun",
            "maintenanceWindowHourUtc": 3,
            "metricsEnabled": true,
            "sqlApiEnabled": true,
            "sqlApiAllowedStatements": [],
            "sqlApiMaxRows": 10000,
            "sqlApiMaxBytes": 10485760,
            "sqlApiTimeoutSeconds": 30,
            "error": "string"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.List()
		if err != nil {
			t.Errorf("Method List failed: %v", err)
		}
	})

	t.Run("Test Create", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Create("<DATABASE_ID>", "<NAME>")
		if err != nil {
			t.Errorf("Method Create failed: %v", err)
		}
	})

	t.Run("Test ListSpecifications", func(t *testing.T) {
		mockResponse := `
{
    "specifications": [
        {
            "slug": "s-2vcpu-2gb",
            "name": "Standard",
            "price": 20,
            "cpu": 2000,
            "memory": 2048,
            "maxConnections": 200,
            "includedStorage": 25,
            "includedBandwidth": 200,
            "enabled": true
        }
    ],
    "total": 9,
    "pricing": {
        "storageOverageRate": 0.125,
        "bandwidthOverageRate": 0.08,
        "replicaRate": 1,
        "pitrRate": 0.2
    }
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListSpecifications()
		if err != nil {
			t.Errorf("Method ListSpecifications failed: %v", err)
		}
	})

	t.Run("Test Get", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Get("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method Get failed: %v", err)
		}
	})

	t.Run("Test Update", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Update("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method Update failed: %v", err)
		}
	})

	t.Run("Test Delete", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.Delete("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method Delete failed: %v", err)
		}
	})

	t.Run("Test ListBackups", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "backups": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "databaseId": "5e5ea5c16897e",
            "projectId": "5e5ea5c16897e",
            "policyId": "5e5ea5c16897e",
            "trigger": "schedule",
            "type": "full",
            "requestedType": "incremental",
            "fallbackReason": "PostgreSQL incremental backups are not offered because they cannot be restored: archived WAL is physical and cannot replay onto a logically-restored base. A full backup was taken instead; use a point-in-time restore (targetTime) to recover to a moment between fulls.",
            "status": "completed",
            "sizeBytes": 1073741824,
            "error": "string"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListBackups("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method ListBackups failed: %v", err)
		}
	})

	t.Run("Test CreateBackup", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "databaseId": "5e5ea5c16897e",
    "projectId": "5e5ea5c16897e",
    "policyId": "5e5ea5c16897e",
    "trigger": "schedule",
    "type": "full",
    "requestedType": "incremental",
    "fallbackReason": "PostgreSQL incremental backups are not offered because they cannot be restored: archived WAL is physical and cannot replay onto a logically-restored base. A full backup was taken instead; use a point-in-time restore (targetTime) to recover to a moment between fulls.",
    "status": "completed",
    "sizeBytes": 1073741824,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateBackup("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method CreateBackup failed: %v", err)
		}
	})

	t.Run("Test ListBackupPolicies", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "policies": [
        {
            "$id": "5e5ea5c16897e",
            "name": "Hourly backups",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "$updatedAt": "2020-10-15T06:38:00.000+00:00",
            "services": [],
            "resources": [],
            "retention": 7,
            "schedule": "0 * * * *",
            "type": "full",
            "enabled": true
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListBackupPolicies("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method ListBackupPolicies failed: %v", err)
		}
	})

	t.Run("Test CreateBackupPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "Hourly backups",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "services": [],
    "resources": [],
    "retention": 7,
    "schedule": "0 * * * *",
    "type": "full",
    "enabled": true
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateBackupPolicy("<DATABASE_ID>", "<POLICY_ID>", "<NAME>", "", 1)
		if err != nil {
			t.Errorf("Method CreateBackupPolicy failed: %v", err)
		}
	})

	t.Run("Test GetBackupPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "Hourly backups",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "services": [],
    "resources": [],
    "retention": 7,
    "schedule": "0 * * * *",
    "type": "full",
    "enabled": true
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetBackupPolicy("<DATABASE_ID>", "<POLICY_ID>")
		if err != nil {
			t.Errorf("Method GetBackupPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdateBackupPolicy", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "name": "Hourly backups",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "services": [],
    "resources": [],
    "retention": 7,
    "schedule": "0 * * * *",
    "type": "full",
    "enabled": true
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateBackupPolicy("<DATABASE_ID>", "<POLICY_ID>")
		if err != nil {
			t.Errorf("Method UpdateBackupPolicy failed: %v", err)
		}
	})

	t.Run("Test DeleteBackupPolicy", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.DeleteBackupPolicy("<DATABASE_ID>", "<POLICY_ID>")
		if err != nil {
			t.Errorf("Method DeleteBackupPolicy failed: %v", err)
		}
	})

	t.Run("Test UpdateBackupStorage", func(t *testing.T) {
		mockResponse := `
{
    "provider": "s3",
    "bucket": "my-backup-bucket",
    "region": "us-east-1",
    "prefix": "backups/",
    "endpoint": "https://minio.example.com"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Expected method PUT, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateBackupStorage("<DATABASE_ID>", "s3", "<BUCKET>", "<ACCESS_KEY>", "<SECRET_KEY>")
		if err != nil {
			t.Errorf("Method UpdateBackupStorage failed: %v", err)
		}
	})

	t.Run("Test GetBackup", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "databaseId": "5e5ea5c16897e",
    "projectId": "5e5ea5c16897e",
    "policyId": "5e5ea5c16897e",
    "trigger": "schedule",
    "type": "full",
    "requestedType": "incremental",
    "fallbackReason": "PostgreSQL incremental backups are not offered because they cannot be restored: archived WAL is physical and cannot replay onto a logically-restored base. A full backup was taken instead; use a point-in-time restore (targetTime) to recover to a moment between fulls.",
    "status": "completed",
    "sizeBytes": 1073741824,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetBackup("<DATABASE_ID>", "<BACKUP_ID>")
		if err != nil {
			t.Errorf("Method GetBackup failed: %v", err)
		}
	})

	t.Run("Test DeleteBackup", func(t *testing.T) {
		mockResponse := `
{
    "message": "success"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.DeleteBackup("<DATABASE_ID>", "<BACKUP_ID>")
		if err != nil {
			t.Errorf("Method DeleteBackup failed: %v", err)
		}
	})

	t.Run("Test ListBranches", func(t *testing.T) {
		mockResponse := `
{
    "branches": [
        {
            "branchId": "a1b2c3d4",
            "branchName": "branch-a1b2c3d4",
            "namespace": "db-myproject-mydb-branch-a1b2c3d4",
            "expiresAt": 1711411200,
            "host": "db-myproject-mydb-a1b2c3d4.fra.appwrite.center",
            "port": 5432,
            "database": "db-myproject-mydb-a1b2c3d4",
            "username": "appwrite",
            "password": "********",
            "ssl": true,
            "engine": "postgresql",
            "connectionString": "postgresql://appwrite:****@db-myproject-mydb-a1b2c3d4.fra.appwrite.center:5432/db-myproject-mydb-a1b2c3d4?sslmode=disable"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListBranches("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method ListBranches failed: %v", err)
		}
	})

	t.Run("Test CreateBranch", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateBranch("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method CreateBranch failed: %v", err)
		}
	})

	t.Run("Test DeleteBranch", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" {
				t.Errorf("Expected method DELETE, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.DeleteBranch("<DATABASE_ID>", "<BRANCH_ID>")
		if err != nil {
			t.Errorf("Method DeleteBranch failed: %v", err)
		}
	})

	t.Run("Test UpdateCredentials", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateCredentials("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method UpdateCredentials failed: %v", err)
		}
	})

	t.Run("Test CreateExecution", func(t *testing.T) {
		mockResponse := `
{
    "rows": [],
    "rowCount": 1,
    "columns": [
        {
            "name": "id",
            "type": "int4"
        }
    ],
    "durationMs": 12,
    "truncated": true,
    "bytes": 1024
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateExecution("<DATABASE_ID>", "<SQL>")
		if err != nil {
			t.Errorf("Method CreateExecution failed: %v", err)
		}
	})

	t.Run("Test CreateFailover", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateFailover("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method CreateFailover failed: %v", err)
		}
	})

	t.Run("Test UpdateMaintenance", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdateMaintenance("<DATABASE_ID>", "sun", 1)
		if err != nil {
			t.Errorf("Method UpdateMaintenance failed: %v", err)
		}
	})

	t.Run("Test CreateMigration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateMigration("<DATABASE_ID>", "shared")
		if err != nil {
			t.Errorf("Method CreateMigration failed: %v", err)
		}
	})

	t.Run("Test ListOperations", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "operations": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "databaseId": "5e5ea5c16897e",
            "type": "update",
            "status": "completed",
            "attempts": 1,
            "errorCode": "Interrupted",
            "errorMessage": "string"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListOperations("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method ListOperations failed: %v", err)
		}
	})

	t.Run("Test GetPitr", func(t *testing.T) {
		mockResponse := `
{
    "earliest": "2020-10-15T06:38:00.000+00:00",
    "latest": "2020-10-15T06:38:00.000+00:00"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetPitr("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method GetPitr failed: %v", err)
		}
	})

	t.Run("Test GetPooler", func(t *testing.T) {
		mockResponse := `
{
    "enabled": true,
    "mode": "transaction",
    "maxConnections": 200,
    "defaultPoolSize": 25,
    "port": 6432,
    "readWriteSplitting": true,
    "poolerCpuRequest": "100m",
    "poolerCpuLimit": "200m",
    "poolerMemoryRequest": "64Mi",
    "poolerMemoryLimit": "128Mi"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetPooler("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method GetPooler failed: %v", err)
		}
	})

	t.Run("Test UpdatePooler", func(t *testing.T) {
		mockResponse := `
{
    "enabled": true,
    "mode": "transaction",
    "maxConnections": 200,
    "defaultPoolSize": 25,
    "port": 6432,
    "readWriteSplitting": true,
    "poolerCpuRequest": "100m",
    "poolerCpuLimit": "200m",
    "poolerMemoryRequest": "64Mi",
    "poolerMemoryLimit": "128Mi"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PATCH" {
				t.Errorf("Expected method PATCH, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.UpdatePooler("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method UpdatePooler failed: %v", err)
		}
	})

	t.Run("Test GetReplicas", func(t *testing.T) {
		mockResponse := `
{
    "replicas": 2,
    "syncMode": "async",
    "syncDegraded": true,
    "syncAcknowledgements": 1,
    "syncStandbyCount": 2,
    "members": [
        {
            "$id": "1",
            "role": "replica",
            "status": "active"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetReplicas("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method GetReplicas failed: %v", err)
		}
	})

	t.Run("Test ListRestorations", func(t *testing.T) {
		mockResponse := `
{
    "total": 5,
    "restorations": [
        {
            "$id": "5e5ea5c16897e",
            "$createdAt": "2020-10-15T06:38:00.000+00:00",
            "databaseId": "5e5ea5c16897e",
            "sourceDatabaseId": "5e5ea5c16897e",
            "projectId": "5e5ea5c16897e",
            "backupId": "5e5ea5c16897e",
            "type": "backup",
            "status": "completed",
            "targetTime": "2020-10-15T06:38:00.000+00:00",
            "startedAt": "2020-10-15T06:38:00.000+00:00",
            "completedAt": "2020-10-15T06:38:00.000+00:00",
            "error": "string"
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.ListRestorations("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method ListRestorations failed: %v", err)
		}
	})

	t.Run("Test CreateRestoration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "databaseId": "5e5ea5c16897e",
    "sourceDatabaseId": "5e5ea5c16897e",
    "projectId": "5e5ea5c16897e",
    "backupId": "5e5ea5c16897e",
    "type": "backup",
    "status": "completed",
    "targetTime": "2020-10-15T06:38:00.000+00:00",
    "startedAt": "2020-10-15T06:38:00.000+00:00",
    "completedAt": "2020-10-15T06:38:00.000+00:00",
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateRestoration("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method CreateRestoration failed: %v", err)
		}
	})

	t.Run("Test GetRestoration", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "databaseId": "5e5ea5c16897e",
    "sourceDatabaseId": "5e5ea5c16897e",
    "projectId": "5e5ea5c16897e",
    "backupId": "5e5ea5c16897e",
    "type": "backup",
    "status": "completed",
    "targetTime": "2020-10-15T06:38:00.000+00:00",
    "startedAt": "2020-10-15T06:38:00.000+00:00",
    "completedAt": "2020-10-15T06:38:00.000+00:00",
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetRestoration("<DATABASE_ID>", "<RESTORATION_ID>")
		if err != nil {
			t.Errorf("Method GetRestoration failed: %v", err)
		}
	})

	t.Run("Test GetStatus", func(t *testing.T) {
		mockResponse := `
{
    "health": "healthy",
    "ready": true,
    "engine": "postgresql",
    "version": "17",
    "uptime": 86400,
    "connections": {
        "current": 12,
        "max": 100
    },
    "syncMode": "async",
    "syncDegraded": true,
    "syncAcknowledgements": 1,
    "syncStandbyCount": 2,
    "replicas": [
        {
            "index": 0,
            "role": "primary",
            "healthy": true
        }
    ],
    "volumes": [
        {
            "path": "/var/lib/postgresql/data",
            "usedPercent": "45%",
            "available": "55GB",
            "mounted": true
        }
    ]
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				t.Errorf("Expected method GET, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.GetStatus("<DATABASE_ID>")
		if err != nil {
			t.Errorf("Method GetStatus failed: %v", err)
		}
	})

	t.Run("Test CreateUpgrade", func(t *testing.T) {
		mockResponse := `
{
    "$id": "5e5ea5c16897e",
    "$createdAt": "2020-10-15T06:38:00.000+00:00",
    "$updatedAt": "2020-10-15T06:38:00.000+00:00",
    "projectId": "5e5ea5c16897e",
    "name": "My Production Database",
    "api": "postgresql",
    "engine": "postgresql",
    "version": "16",
    "specification": "s-2vcpu-2gb",
    "backend": "edge",
    "hostname": "db-myproject-mydb.fra.appwrite.center",
    "connectionPort": 5432,
    "connectionUser": "appwrite_user",
    "connectionPassword": "••••••••",
    "connectionString": "postgresql://user:pass@db-myproject-mydb.fra.appwrite.center:5432/postgres?sslmode=require",
    "ssl": true,
    "status": "ready",
    "containerStatus": "active",
    "lifecycleState": "active",
    "idleTimeoutMinutes": 15,
    "cpu": 2000,
    "memory": 4096,
    "storage": 100,
    "storageClass": "ssd",
    "storageMaxGb": 100,
    "nodePool": "db-pool-4vcpu-8gb",
    "replicas": 2,
    "syncMode": "async",
    "networkMaxConnections": 500,
    "networkIdleTimeoutSeconds": 900,
    "networkIPAllowlist": [],
    "backupEnabled": true,
    "pitr": true,
    "pitrRetentionDays": 14,
    "storageAutoscaling": true,
    "storageAutoscalingThresholdPercent": 85,
    "storageAutoscalingMaxGb": 500,
    "maintenanceWindowDay": "sun",
    "maintenanceWindowHourUtc": 3,
    "metricsEnabled": true,
    "sqlApiEnabled": true,
    "sqlApiAllowedStatements": [],
    "sqlApiMaxRows": 10000,
    "sqlApiMaxBytes": 10485760,
    "sqlApiTimeoutSeconds": 30,
    "error": "string"
}
`

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Expected method POST, got %s", r.Method)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(mockResponse))
		}))
		defer ts.Close()

		srv := New(newTestClient(ts))

		_, err := srv.CreateUpgrade("<DATABASE_ID>", "<TARGET_VERSION>")
		if err != nil {
			t.Errorf("Method CreateUpgrade failed: %v", err)
		}
	})
}
