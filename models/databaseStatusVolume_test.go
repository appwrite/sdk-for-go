package models

import (
	"encoding/json"
	"reflect"
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

func TestDatabaseStatusVolumeModelDefaults(t *testing.T) {
	var fresh DatabaseStatusVolume
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Path)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing path to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"0%\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.UsedPercent)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing usedPercent to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"0GB\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Available)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing available to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Mounted)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing mounted to default to %v, got %v", want, got)
		}
	}

	reused := DatabaseStatusVolume{Path: "/var/lib/postgresql/data", UsedPercent: "45%", Available: "55GB", Mounted: true}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
