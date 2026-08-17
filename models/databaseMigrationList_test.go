package models

import (
	"encoding/json"
	"testing"
)

func TestDatabaseMigrationListModel(t *testing.T) {
	model := DatabaseMigrationList{Total: 5, Migrations: []DatabaseMigration{DatabaseMigration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", ProjectId: "5e5ea5c16897e", DatabaseId: "5e5ea5c16897e", Specification: "s-2vcpu-4gb", Phase: "pending", Attempt: 0, LastError: "string", LagDocuments: 0, VerifiedAt: "2020-10-15T06:38:00.000+00:00", CutoverAt: "2020-10-15T06:38:00.000+00:00", SoakUntil: "2020-10-15T06:38:00.000+00:00", AutoCutover: true, CutoverRequested: true, Paused: true}}}

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
