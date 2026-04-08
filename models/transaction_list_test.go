package models

import (
    "encoding/json"
    "testing"
)

func TestTransactionListModel(t *testing.T) {
    model := TransactionList{        Total: 5,        Transactions: []Transaction{Transaction{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Status: "pending",        Operations: 5,        ExpiresAt: "2020-10-15T06:38:00.000+00:00",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result TransactionList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
