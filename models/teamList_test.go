package models

import (
    "encoding/json"
    "testing"
)

func TestTeamListModel(t *testing.T) {
    model := TeamList{        Total: 5,        Teams: []Team{Team{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "VIP",        Total: 7,        Prefs: Preferences{    },    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result TeamList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
