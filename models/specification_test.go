package models

import (
    "encoding/json"
    "testing"
)

func TestSpecificationModel(t *testing.T) {
    model := Specification{        Memory: 512,        Cpus: 1,        Enabled: true,        Slug: "s-1vcpu-512mb",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Specification
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Memory != model.Memory {
        t.Errorf("Expected Memory %v, got %v", model.Memory, result.Memory)
    }
    if result.Cpus != model.Cpus {
        t.Errorf("Expected Cpus %v, got %v", model.Cpus, result.Cpus)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }
    if result.Slug != model.Slug {
        t.Errorf("Expected Slug %v, got %v", model.Slug, result.Slug)
    }}
