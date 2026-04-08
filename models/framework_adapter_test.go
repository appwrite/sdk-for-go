package models

import (
    "encoding/json"
    "testing"
)

func TestFrameworkAdapterModel(t *testing.T) {
    model := FrameworkAdapter{        Key: "static",        InstallCommand: "npm install",        BuildCommand: "npm run build",        OutputDirectory: "./dist",        FallbackFile: "index.html",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result FrameworkAdapter
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Key != model.Key {
        t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
    }
    if result.InstallCommand != model.InstallCommand {
        t.Errorf("Expected InstallCommand %v, got %v", model.InstallCommand, result.InstallCommand)
    }
    if result.BuildCommand != model.BuildCommand {
        t.Errorf("Expected BuildCommand %v, got %v", model.BuildCommand, result.BuildCommand)
    }
    if result.OutputDirectory != model.OutputDirectory {
        t.Errorf("Expected OutputDirectory %v, got %v", model.OutputDirectory, result.OutputDirectory)
    }
    if result.FallbackFile != model.FallbackFile {
        t.Errorf("Expected FallbackFile %v, got %v", model.FallbackFile, result.FallbackFile)
    }}
