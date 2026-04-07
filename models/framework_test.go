package models

import (
    "encoding/json"
    "testing"
)

func TestFrameworkModel(t *testing.T) {
    model := Framework{        Key: "sveltekit",        Name: "SvelteKit",        BuildRuntime: "node-22",        Runtimes: []string{"test"},        Adapters: []FrameworkAdapter{FrameworkAdapter{        Key: "static",        InstallCommand: "npm install",        BuildCommand: "npm run build",        OutputDirectory: "./dist",        FallbackFile: "index.html",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Framework
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Key != model.Key {
        t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.BuildRuntime != model.BuildRuntime {
        t.Errorf("Expected BuildRuntime %v, got %v", model.BuildRuntime, result.BuildRuntime)
    }}
