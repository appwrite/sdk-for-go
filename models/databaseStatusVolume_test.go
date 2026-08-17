package models

import (
	"encoding/json"
	"testing"
)

func TestDatabaseStatusVolumeModel(t *testing.T) {
	model := DatabaseStatusVolume{Path: "/var/lib/postgresql/data", UsedPercent: "45%", Available: "55GB", Mounted: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DatabaseStatusVolume
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Path != model.Path {
		t.Errorf("Expected Path %v, got %v", model.Path, result.Path)
	}
	if result.UsedPercent != model.UsedPercent {
		t.Errorf("Expected UsedPercent %v, got %v", model.UsedPercent, result.UsedPercent)
	}
	if result.Available != model.Available {
		t.Errorf("Expected Available %v, got %v", model.Available, result.Available)
	}
	if result.Mounted != model.Mounted {
		t.Errorf("Expected Mounted %v, got %v", model.Mounted, result.Mounted)
	}
}
