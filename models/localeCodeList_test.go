package models

import (
    "encoding/json"
    "testing"
)

func TestLocaleCodeListModel(t *testing.T) {
    model := LocaleCodeList{        Total: 5,        LocaleCodes: []LocaleCode{LocaleCode{        Code: "en-us",        Name: "US",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result LocaleCodeList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
