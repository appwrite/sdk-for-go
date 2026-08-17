package models

import (
	"encoding/json"
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
