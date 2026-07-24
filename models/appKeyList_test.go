package models

import (
    "encoding/json"
    "testing"
)

func TestAppKeyListModel(t *testing.T) {
    model := AppKeyList{        Total: 5,        Keys: []AppKey{AppKey{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        AppId: "5e5ea5c16897e",        Secret: "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a",        Hint: "f5c6c7",        CreatedById: "5e5ea5c16897e",        CreatedByName: "Walter White",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AppKeyList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
