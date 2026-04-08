package models

import (
    "encoding/json"
    "testing"
)

func TestColumnIndexListModel(t *testing.T) {
    model := ColumnIndexList{        Total: 5,        Indexes: []ColumnIndex{ColumnIndex{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "index1",        Type: "primary",        Status: "available",        Error: "string",        Columns: []string{"test"},        Lengths: []int{1},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ColumnIndexList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
