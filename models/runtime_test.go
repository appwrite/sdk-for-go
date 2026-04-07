package models

import (
    "encoding/json"
    "testing"
)

func TestRuntimeModel(t *testing.T) {
    model := Runtime{        Id: "python-3.8",        Key: "python",        Name: "Python",        Version: "3.8",        Base: "python:3.8-alpine",        Image: "appwrite\\/runtime-for-python:3.8",        Logo: "python.png",        Supports: []string{"test"},    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Runtime
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.Key != model.Key {
        t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Version != model.Version {
        t.Errorf("Expected Version %v, got %v", model.Version, result.Version)
    }
    if result.Base != model.Base {
        t.Errorf("Expected Base %v, got %v", model.Base, result.Base)
    }
    if result.Image != model.Image {
        t.Errorf("Expected Image %v, got %v", model.Image, result.Image)
    }
    if result.Logo != model.Logo {
        t.Errorf("Expected Logo %v, got %v", model.Logo, result.Logo)
    }}
