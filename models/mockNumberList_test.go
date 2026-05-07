package models

import (
    "encoding/json"
    "testing"
)

func TestMockNumberListModel(t *testing.T) {
    model := MockNumberList{        Total: 5,        MockNumbers: []MockNumber{MockNumber{        Number: "+1612842323",        Otp: "123456",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result MockNumberList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
