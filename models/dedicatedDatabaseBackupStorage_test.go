package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseBackupStorageModel(t *testing.T) {
	model := DedicatedDatabaseBackupStorage{Provider: "s3", Bucket: "my-backup-bucket", Region: "us-east-1", Prefix: "backups/", Endpoint: "https://minio.example.com"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseBackupStorage
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Provider != model.Provider {
		t.Errorf("Expected Provider %v, got %v", model.Provider, result.Provider)
	}
	if result.Bucket != model.Bucket {
		t.Errorf("Expected Bucket %v, got %v", model.Bucket, result.Bucket)
	}
	if result.Region != model.Region {
		t.Errorf("Expected Region %v, got %v", model.Region, result.Region)
	}
	if result.Prefix != model.Prefix {
		t.Errorf("Expected Prefix %v, got %v", model.Prefix, result.Prefix)
	}
	if result.Endpoint != model.Endpoint {
		t.Errorf("Expected Endpoint %v, got %v", model.Endpoint, result.Endpoint)
	}
}
