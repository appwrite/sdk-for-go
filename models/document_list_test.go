package models

import (
    "encoding/json"
    "testing"
)

func TestDocumentListModel(t *testing.T) {
    model := DocumentList{        Total: 5,        Documents: []Document{Document{        Id: "5e5ea5c16897e",        Sequence: "1",        CollectionId: "5e5ea5c15117e",        DatabaseId: "5e5ea5c15117e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result DocumentList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
