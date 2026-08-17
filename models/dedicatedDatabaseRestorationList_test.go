package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseRestorationListModel(t *testing.T) {
	model := DedicatedDatabaseRestorationList{Total: 5, Restorations: []DedicatedDatabaseRestoration{DedicatedDatabaseRestoration{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", DatabaseId: "5e5ea5c16897e", SourceDatabaseId: "5e5ea5c16897e", ProjectId: "5e5ea5c16897e", BackupId: "5e5ea5c16897e", Type: "backup", Status: "completed", TargetTime: "2020-10-15T06:38:00.000+00:00", StartedAt: "2020-10-15T06:38:00.000+00:00", CompletedAt: "2020-10-15T06:38:00.000+00:00", Error: "string"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseRestorationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
