package models

import (
    "encoding/json"
    "testing"
)

func TestRowListModel(t *testing.T) {
    model := RowList{        Total: 5,        Rows: []Row{Row{        Id: "5e5ea5c16897e",        Sequence: "1",        TableId: "5e5ea5c15117e",        DatabaseId: "5e5ea5c15117e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result RowList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
