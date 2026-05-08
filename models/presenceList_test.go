package models

import (
    "encoding/json"
    "testing"
)

func TestPresenceListModel(t *testing.T) {
    model := PresenceList{        Total: 5,        Presences: []Presence{Presence{        Id: "5e5ea5c16897e",        Sequence: "1",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Permissions: []string{"test"},        UserInternalId: "1",        UserId: "674af8f3e12a5f9ac0be",        Source: "HTTP",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PresenceList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
