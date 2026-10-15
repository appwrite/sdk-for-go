package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDedicatedDatabaseExecutionModel(t *testing.T) {
	model := DedicatedDatabaseExecution{Rows: []interface{}{}, RowCount: 1, Columns: []DedicatedDatabaseExecutionColumn{DedicatedDatabaseExecutionColumn{Name: "id", Type: "int4"}}, DurationMs: 12, Truncated: true, Bytes: 1024}

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

func TestDedicatedDatabaseExecutionModelDefaults(t *testing.T) {
	var fresh DedicatedDatabaseExecution
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Rows)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing rows to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RowCount)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing rowCount to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DurationMs)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing durationMs to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Truncated)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing truncated to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing bytes to default to %v, got %v", want, got)
		}
	}

	reused := DedicatedDatabaseExecution{Rows: []interface{}{}, RowCount: 1, Columns: []DedicatedDatabaseExecutionColumn{DedicatedDatabaseExecutionColumn{Name: "id", Type: "int4"}}, DurationMs: 12, Truncated: true, Bytes: 1024}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
