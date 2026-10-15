package models

import (
	"encoding/json"
	"reflect"
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

func TestDedicatedDatabasePITRWindowsModelDefaults(t *testing.T) {
	var fresh DedicatedDatabasePITRWindows
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Earliest)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing earliest to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Latest)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing latest to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabasePITRWindows{Earliest: "2020-10-15T06:38:00.000+00:00", Latest: "2020-10-15T06:38:00.000+00:00"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
