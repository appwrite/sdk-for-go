package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDatabaseMigrationListModel(t *testing.T) {
	model := DatabaseMigrationList{Total: 5, Migrations: []DatabaseMigration{DatabaseMigration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", DatabaseId: "5e5ea5c16897e", Specification: "s-2vcpu-4gb", Phase: "pending", Attempt: 0, LastError: "string", LagDocuments: 0, ChangelogWatermark: 0, VerifiedAt: "2020-10-15T06:38:00.000+00:00", CutoverAt: "2020-10-15T06:38:00.000+00:00", SoakUntil: "2020-10-15T06:38:00.000+00:00", AutoCutover: true, CutoverRequested: true, Paused: true}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DatabaseMigrationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestDatabaseMigrationListModelDefaults(t *testing.T) {
	var fresh DatabaseMigrationList
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

	reused := DatabaseMigrationList{Total: 5, Migrations: []DatabaseMigration{DatabaseMigration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", DatabaseId: "5e5ea5c16897e", Specification: "s-2vcpu-4gb", Phase: "pending", Attempt: 0, LastError: "string", LagDocuments: 0, ChangelogWatermark: 0, VerifiedAt: "2020-10-15T06:38:00.000+00:00", CutoverAt: "2020-10-15T06:38:00.000+00:00", SoakUntil: "2020-10-15T06:38:00.000+00:00", AutoCutover: true, CutoverRequested: true, Paused: true}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
