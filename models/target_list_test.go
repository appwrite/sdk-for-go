package models

import (
    "encoding/json"
    "testing"
)

func TestTargetListModel(t *testing.T) {
    model := TargetList{        Total: 5,        Targets: []Target{Target{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Apple iPhone 12",        UserId: "259125845563242502",        ProviderType: "email",        Identifier: "token",        Expired: true,    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result TargetList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
