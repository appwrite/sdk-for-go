package models

import (
	"encoding/json"
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
