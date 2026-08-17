package models

import (
	"encoding/json"
	"testing"
)

func TestDedicatedDatabaseExecutionModel(t *testing.T) {
	model := DedicatedDatabaseExecution{Rows: map[string]interface{}{}, RowCount: 1, Columns: []DedicatedDatabaseExecutionColumn{DedicatedDatabaseExecutionColumn{Name: "id", Type: "int4"}}, DurationMs: 12, Truncated: true, Bytes: 1024}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DedicatedDatabaseExecution
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.RowCount != model.RowCount {
		t.Errorf("Expected RowCount %v, got %v", model.RowCount, result.RowCount)
	}
	if result.DurationMs != model.DurationMs {
		t.Errorf("Expected DurationMs %v, got %v", model.DurationMs, result.DurationMs)
	}
	if result.Truncated != model.Truncated {
		t.Errorf("Expected Truncated %v, got %v", model.Truncated, result.Truncated)
	}
	if result.Bytes != model.Bytes {
		t.Errorf("Expected Bytes %v, got %v", model.Bytes, result.Bytes)
	}
}
