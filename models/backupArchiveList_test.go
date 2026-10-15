package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBackupArchiveListModel(t *testing.T) {
	model := BackupArchiveList{Total: 5, Archives: []BackupArchive{BackupArchive{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", PolicyId: "did8jx6ws45jana098ab7", Status: "completed", MigrationId: "did8jx6ws45jana098ab7", Services: []string{"test"}, Resources: []string{"test"}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result BackupArchiveList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestBackupArchiveListModelDefaults(t *testing.T) {
	var fresh BackupArchiveList
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

	reused := BackupArchiveList{Total: 5, Archives: []BackupArchive{BackupArchive{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", PolicyId: "did8jx6ws45jana098ab7", Status: "completed", MigrationId: "did8jx6ws45jana098ab7", Services: []string{"test"}, Resources: []string{"test"}}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
