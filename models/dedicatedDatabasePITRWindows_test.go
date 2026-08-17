package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabasePITRWindowsModel(t *testing.T) {
	model := DedicatedDatabasePITRWindows{Earliest: "2020-10-15T06:38:00.000+00:00", Latest: "2020-10-15T06:38:00.000+00:00"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabasePITRWindows
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Earliest != model.Earliest {
		t.Errorf("Expected Earliest %v, got %v", model.Earliest, result.Earliest)
	}
	if result.Latest != model.Latest {
		t.Errorf("Expected Latest %v, got %v", model.Latest, result.Latest)
	}
}
