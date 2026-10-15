package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseBackupListModel(t *testing.T) {
	model := DedicatedDatabaseBackupList{Total: 5, Backups: []DedicatedDatabaseBackup{DedicatedDatabaseBackup{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", DatabaseId: "5e5ea5c16897e", ProjectId: "5e5ea5c16897e", PolicyId: "5e5ea5c16897e", Trigger: "schedule", Type: "full", RequestedType: "incremental", FallbackReason: "PostgreSQL incremental backups are not offered because they cannot be restored: archived WAL is physical and cannot replay onto a logically-restored base. A full backup was taken instead; use a point-in-time restore (targetTime) to recover to a moment between fulls.", Status: "completed", SizeBytes: 1073741824, Error: "string"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseBackupList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestDedicatedDatabaseBackupListModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseBackupList
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

	reused := DedicatedDatabaseBackupList{Total: 5, Backups: []DedicatedDatabaseBackup{DedicatedDatabaseBackup{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", DatabaseId: "5e5ea5c16897e", ProjectId: "5e5ea5c16897e", PolicyId: "5e5ea5c16897e", Trigger: "schedule", Type: "full", RequestedType: "incremental", FallbackReason: "PostgreSQL incremental backups are not offered because they cannot be restored: archived WAL is physical and cannot replay onto a logically-restored base. A full backup was taken instead; use a point-in-time restore (targetTime) to recover to a moment between fulls.", Status: "completed", SizeBytes: 1073741824, Error: "string"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
