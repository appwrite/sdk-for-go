package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseExecutionColumnModel(t *testing.T) {
	model := DedicatedDatabaseExecutionColumn{Name: "id", Type: "int4"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseExecutionColumn
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
}
