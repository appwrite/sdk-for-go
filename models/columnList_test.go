package models

import (
    "encoding/json"
    "testing"
)

func TestColumnListModel(t *testing.T) {
    model := ColumnList{        Total: 5,        Columns: []interface{}{},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ColumnList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
