package models

import (
	"encoding/json"
	"reflect"
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

func TestDedicatedDatabaseBackupStorageModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseBackupStorage
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Provider)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing provider to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Bucket)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing bucket to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Region)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing region to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"backups/\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Prefix)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing prefix to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Endpoint)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing endpoint to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabaseBackupStorage{Provider: "s3", Bucket: "my-backup-bucket", Region: "us-east-1", Prefix: "backups/", Endpoint: "https://minio.example.com"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
