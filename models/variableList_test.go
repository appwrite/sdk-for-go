package models

import (
    "encoding/json"
    "testing"
)

func TestVariableListModel(t *testing.T) {
    model := VariableList{        Total: 5,        Variables: []Variable{Variable{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "API_KEY",        Value: "myPa$$word1",        Secret: true,        ResourceType: "function",        ResourceId: "myAwesomeFunction",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result VariableList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
