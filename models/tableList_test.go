package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTableListModel(t *testing.T) {
	model := TableList{Total: 5, Tables: []Table{Table{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, DatabaseId: "5e5ea5c16897e", Name: "My Table", Description: "Movies available in the catalog.", Enabled: true, RowSecurity: true, Columns: []interface{}{}, Indexes: []ColumnIndex{ColumnIndex{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "index1", Description: "Speeds up lookups by email.", Type: "primary", Status: "available", Error: "string", Columns: []string{"test"}, Lengths: []int{1}}}, BytesMax: 65535, BytesUsed: 1500}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result TableList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestTableListModelDefaults(t *testing.T) {
	var fresh TableList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := TableList{Total: 5, Tables: []Table{Table{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Permissions: []string{"test"}, DatabaseId: "5e5ea5c16897e", Name: "My Table", Description: "Movies available in the catalog.", Enabled: true, RowSecurity: true, Columns: []interface{}{}, Indexes: []ColumnIndex{ColumnIndex{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Key: "index1", Description: "Speeds up lookups by email.", Type: "primary", Status: "available", Error: "string", Columns: []string{"test"}, Lengths: []int{1}}}, BytesMax: 65535, BytesUsed: 1500}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
