package models

import (
    "encoding/json"
    "testing"
)

func TestPhoneListModel(t *testing.T) {
    model := PhoneList{        Total: 5,        Phones: []Phone{Phone{        Code: "+1",        CountryCode: "US",        CountryName: "United States",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result PhoneList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
