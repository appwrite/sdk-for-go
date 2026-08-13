package models

import (
    "encoding/json"
    "testing"
)

func TestDedicatedDatabaseOperationListModel(t *testing.T) {
    model := DedicatedDatabaseOperationList{        Total: 5,        Operations: []DedicatedDatabaseOperation{DedicatedDatabaseOperation{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        DatabaseId: "5e5ea5c16897e",        Type: "update",        Status: "completed",        Attempts: 1,        ErrorCode: "Interrupted",        ErrorMessage: "string",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DedicatedDatabaseOperationList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
