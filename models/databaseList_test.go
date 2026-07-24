package models

import (
    "encoding/json"
    "testing"
)

func TestDatabaseListModel(t *testing.T) {
    model := DatabaseList{        Total: 5,        Databases: []Database{Database{        Id: "5e5ea5c16897e",        Name: "My Database",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Enabled: true,        Type: "legacy",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DatabaseList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
