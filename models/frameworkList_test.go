package models

import (
    "encoding/json"
    "testing"
)

func TestFrameworkListModel(t *testing.T) {
    model := FrameworkList{        Total: 5,        Frameworks: []Framework{Framework{        Key: "sveltekit",        Name: "SvelteKit",        BuildRuntime: "node-22",        Runtimes: []string{"test"},        Adapters: []FrameworkAdapter{FrameworkAdapter{        Key: "static",        InstallCommand: "npm install",        BuildCommand: "npm run build",        OutputDirectory: "./dist",        FallbackFile: "index.html",    },
            },    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result FrameworkList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
