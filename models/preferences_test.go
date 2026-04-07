package models

import (
    "encoding/json"
    "testing"
)

func TestPreferencesModel(t *testing.T) {
    model := Preferences{    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Preferences
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }}
