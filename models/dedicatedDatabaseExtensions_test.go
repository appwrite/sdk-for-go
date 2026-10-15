package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseExtensionsModel(t *testing.T) {
	model := DedicatedDatabaseExtensions{Installed: []string{"test"}, Available: []string{"test"}, Metadata: []PostgresExtension{PostgresExtension{Key: "vector", Name: "pgvector", Description: "Vector data type and similarity search for embeddings.", Category: "search"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseExtensions
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDedicatedDatabaseExtensionsModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseExtensions
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Installed)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing installed to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
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

	reused := DedicatedDatabaseExtensions{Installed: []string{"test"}, Available: []string{"test"}, Metadata: []PostgresExtension{PostgresExtension{Key: "vector", Name: "pgvector", Description: "Vector data type and similarity search for embeddings.", Category: "search"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
