package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBackupPolicyListModel(t *testing.T) {
	model := BackupPolicyList{Total: 5, Policies: []BackupPolicy{BackupPolicy{Id: "5e5ea5c16897e", Name: "Hourly backups", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Services: []string{"test"}, Resources: []string{"test"}, Retention: 7, Schedule: "0 * * * *", Type: "full", Enabled: true}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result BackupPolicyList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestBackupPolicyListModelDefaults(t *testing.T) {
	var fresh BackupPolicyList
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

	reused := BackupPolicyList{Total: 5, Policies: []BackupPolicy{BackupPolicy{Id: "5e5ea5c16897e", Name: "Hourly backups", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Services: []string{"test"}, Resources: []string{"test"}, Retention: 7, Schedule: "0 * * * *", Type: "full", Enabled: true}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
