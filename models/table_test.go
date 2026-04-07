package models

import (
    "encoding/json"
    "testing"
)

func TestTableModel(t *testing.T) {
    model := Table{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        DatabaseId: "5e5ea5c16897e",        Name: "My Table",        Enabled: true,        RowSecurity: true,        Columns: []interface{}{},        Indexes: []ColumnIndex{ColumnIndex{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "index1",        Type: "primary",        Status: "available",        Error: "string",        Columns: []string{"test"},        Lengths: []int{1},    },
            },        BytesMax: 65535,        BytesUsed: 1500,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Table
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
    if result.RowSecurity != model.RowSecurity {
        t.Errorf("Expected RowSecurity %v, got %v", model.RowSecurity, result.RowSecurity)
    }
    if result.BytesMax != model.BytesMax {
        t.Errorf("Expected BytesMax %v, got %v", model.BytesMax, result.BytesMax)
    }
    if result.BytesUsed != model.BytesUsed {
        t.Errorf("Expected BytesUsed %v, got %v", model.BytesUsed, result.BytesUsed)
    }}
