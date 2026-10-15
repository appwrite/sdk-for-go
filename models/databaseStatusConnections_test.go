package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDatabaseStatusConnectionsModel(t *testing.T) {
	model := DatabaseStatusConnections{Current: 12, Max: 100}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DatabaseStatusConnections
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Current != model.Current {
		t.Errorf("Expected Current %v, got %v", model.Current, result.Current)
	}
	if result.Max != model.Max {
		t.Errorf("Expected Max %v, got %v", model.Max, result.Max)
	}
}

func TestDatabaseStatusConnectionsModelDefaults(t *testing.T) {
	var fresh DatabaseStatusConnections
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Current)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing current to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("100"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Max)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing max to default to %v, got %v", want, got)
		}
	}

	reused := DatabaseStatusConnections{Current: 12, Max: 100}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
