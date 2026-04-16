package models

import (
    "encoding/json"
    "testing"
)

func TestKeyListModel(t *testing.T) {
    model := KeyList{        Total: 5,        Keys: []Key{Key{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My API Key",        Expire: "2020-10-15T06:38:00.000+00:00",        Scopes: []string{"test"},        Secret: "919c2d18fb5d4...a2ae413da83346ad2",        AccessedAt: "2020-10-15T06:38:00.000+00:00",        Sdks: []string{"test"},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result KeyList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
