package models

import (
    "encoding/json"
    "testing"
)

func TestSpecificationListModel(t *testing.T) {
    model := SpecificationList{        Total: 5,        Specifications: []Specification{Specification{        Memory: 512,        Cpus: 1,        Enabled: true,        Slug: "s-1vcpu-512mb",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result SpecificationList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
