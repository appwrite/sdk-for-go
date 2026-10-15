package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBackupRestorationListModel(t *testing.T) {
	model := BackupRestorationList{Total: 5, Restorations: []BackupRestoration{BackupRestoration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ArchiveId: "did8jx6ws45jana098ab7", PolicyId: "did8jx6ws45jana098ab7", Status: "completed", MigrationId: "did8jx6ws45jana098ab7", Services: []string{"test"}, Resources: []string{"test"}, Options: map[string]interface{}{}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result BackupRestorationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestBackupRestorationListModelDefaults(t *testing.T) {
	var fresh BackupRestorationList
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

	reused := BackupRestorationList{Total: 5, Restorations: []BackupRestoration{BackupRestoration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ArchiveId: "did8jx6ws45jana098ab7", PolicyId: "did8jx6ws45jana098ab7", Status: "completed", MigrationId: "did8jx6ws45jana098ab7", Services: []string{"test"}, Resources: []string{"test"}, Options: map[string]interface{}{}}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
