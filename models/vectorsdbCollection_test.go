package models

import (
	"encoding/json"
	"testing"
)

func TestVectorsdbCollectionModel(t *testing.T) {
	model := VectorsdbCollection{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, DatabaseId: "5e5ea5c16897e", Name: "My Collection", Enabled: true, DocumentSecurity: true, Attributes: []map[string]any{}, Indexes: []Index{Index{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "index1", Type: "primary", Status: "available", Error: "string", Attributes: []string{"test"}, Lengths: []int{1}}}, BytesMax: 65535, BytesUsed: 1500, Dimension: 1536}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result VectorsdbCollection
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
	if result.DatabaseId != model.DatabaseId {
		t.Errorf("Expected DatabaseId %v, got %v", model.DatabaseId, result.DatabaseId)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
	if result.DocumentSecurity != model.DocumentSecurity {
		t.Errorf("Expected DocumentSecurity %v, got %v", model.DocumentSecurity, result.DocumentSecurity)
	}
	if result.BytesMax != model.BytesMax {
		t.Errorf("Expected BytesMax %v, got %v", model.BytesMax, result.BytesMax)
	}
	if result.BytesUsed != model.BytesUsed {
		t.Errorf("Expected BytesUsed %v, got %v", model.BytesUsed, result.BytesUsed)
	}
	if result.Dimension != model.Dimension {
		t.Errorf("Expected Dimension %v, got %v", model.Dimension, result.Dimension)
	}
}
