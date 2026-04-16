package models

import (
    "encoding/json"
    "testing"
)

func TestRuntimeListModel(t *testing.T) {
    model := RuntimeList{        Total: 5,        Runtimes: []Runtime{Runtime{        Id: "python-3.8",        Key: "python",        Name: "Python",        Version: "3.8",        Base: "python:3.8-alpine",        Image: "appwrite\\/runtime-for-python:3.8",        Logo: "python.png",        Supports: []string{"test"},    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result RuntimeList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
