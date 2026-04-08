package models

import (
    "encoding/json"
    "testing"
)

func TestCollectionListModel(t *testing.T) {
    model := CollectionList{        Total: 5,        Collections: []Collection{Collection{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        DatabaseId: "5e5ea5c16897e",        Name: "My Collection",        Enabled: true,        DocumentSecurity: true,        Attributes: []map[string]any{},        Indexes: []Index{Index{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "index1",        Type: "primary",        Status: "available",        Error: "string",        Attributes: []string{"test"},        Lengths: []int{1},    },
            },        BytesMax: 65535,        BytesUsed: 1500,    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result CollectionList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
