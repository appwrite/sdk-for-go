package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPostgresExtensionModel(t *testing.T) {
	model := PostgresExtension{Key: "vector", Name: "pgvector", Description: "Vector data type and similarity search for embeddings.", Category: "search"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result PostgresExtension
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Key != model.Key {
		t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Description != model.Description {
		t.Errorf("Expected Description %v, got %v", model.Description, result.Description)
	}
	if result.Category != model.Category {
		t.Errorf("Expected Category %v, got %v", model.Category, result.Category)
	}
}

func TestPostgresExtensionModelDefaults(t *testing.T) {
	var fresh PostgresExtension
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Key)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing key to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Name)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing name to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Description)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing description to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Category)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing category to default to %v, got %v", want, got)
		}
	}

	reused := PostgresExtension{Key: "vector", Name: "pgvector", Description: "Vector data type and similarity search for embeddings.", Category: "search"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
