package models

import (
    "encoding/json"
    "testing"
)

func TestSiteModel(t *testing.T) {
    model := Site{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Site",        Enabled: true,        Live: true,        Logging: true,        Framework: "react",        DeploymentRetention: 7,        DeploymentId: "5e5ea5c16897e",        DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00",        DeploymentScreenshotLight: "5e5ea5c16897e",        DeploymentScreenshotDark: "5e5ea5c16897e",        LatestDeploymentId: "5e5ea5c16897e",        LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00",        LatestDeploymentStatus: "ready",        Vars: []Variable{Variable{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "API_KEY",        Value: "myPa$$word1",        Secret: true,        ResourceType: "function",        ResourceId: "myAwesomeFunction",    },
            },        Timeout: 300,        InstallCommand: "npm install",        BuildCommand: "npm run build",        StartCommand: "node custom-server.mjs",        OutputDirectory: "build",        InstallationId: "6m40at4ejk5h2u9s1hboo",        ProviderRepositoryId: "appwrite",        ProviderBranch: "main",        ProviderRootDirectory: "sites/helloWorld",        ProviderSilentMode: true,        BuildSpecification: "s-1vcpu-512mb",        RuntimeSpecification: "s-1vcpu-512mb",        BuildRuntime: "node-22",        Adapter: "static",        FallbackFile: "index.html",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Site
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Id != model.Id {
        t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
    }
    if result.CreatedAt != model.CreatedAt {
        t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
    }
    if result.UpdatedAt != model.UpdatedAt {
        t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
    }
    if result.Name != model.Name {
        t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }
    if result.Live != model.Live {
        t.Errorf("Expected Live %v, got %v", model.Live, result.Live)
    }
    if result.Logging != model.Logging {
        t.Errorf("Expected Logging %v, got %v", model.Logging, result.Logging)
    }
    if result.Framework != model.Framework {
        t.Errorf("Expected Framework %v, got %v", model.Framework, result.Framework)
    }
    if result.DeploymentRetention != model.DeploymentRetention {
        t.Errorf("Expected DeploymentRetention %v, got %v", model.DeploymentRetention, result.DeploymentRetention)
    }
    if result.DeploymentId != model.DeploymentId {
        t.Errorf("Expected DeploymentId %v, got %v", model.DeploymentId, result.DeploymentId)
    }
    if result.DeploymentCreatedAt != model.DeploymentCreatedAt {
        t.Errorf("Expected DeploymentCreatedAt %v, got %v", model.DeploymentCreatedAt, result.DeploymentCreatedAt)
    }
    if result.DeploymentScreenshotLight != model.DeploymentScreenshotLight {
        t.Errorf("Expected DeploymentScreenshotLight %v, got %v", model.DeploymentScreenshotLight, result.DeploymentScreenshotLight)
    }
    if result.DeploymentScreenshotDark != model.DeploymentScreenshotDark {
        t.Errorf("Expected DeploymentScreenshotDark %v, got %v", model.DeploymentScreenshotDark, result.DeploymentScreenshotDark)
    }
    if result.LatestDeploymentId != model.LatestDeploymentId {
        t.Errorf("Expected LatestDeploymentId %v, got %v", model.LatestDeploymentId, result.LatestDeploymentId)
    }
    if result.LatestDeploymentCreatedAt != model.LatestDeploymentCreatedAt {
        t.Errorf("Expected LatestDeploymentCreatedAt %v, got %v", model.LatestDeploymentCreatedAt, result.LatestDeploymentCreatedAt)
    }
    if result.LatestDeploymentStatus != model.LatestDeploymentStatus {
        t.Errorf("Expected LatestDeploymentStatus %v, got %v", model.LatestDeploymentStatus, result.LatestDeploymentStatus)
    }
    if result.Timeout != model.Timeout {
        t.Errorf("Expected Timeout %v, got %v", model.Timeout, result.Timeout)
    }
    if result.InstallCommand != model.InstallCommand {
        t.Errorf("Expected InstallCommand %v, got %v", model.InstallCommand, result.InstallCommand)
    }
    if result.BuildCommand != model.BuildCommand {
        t.Errorf("Expected BuildCommand %v, got %v", model.BuildCommand, result.BuildCommand)
    }
    if result.StartCommand != model.StartCommand {
        t.Errorf("Expected StartCommand %v, got %v", model.StartCommand, result.StartCommand)
    }
    if result.OutputDirectory != model.OutputDirectory {
        t.Errorf("Expected OutputDirectory %v, got %v", model.OutputDirectory, result.OutputDirectory)
    }
    if result.InstallationId != model.InstallationId {
        t.Errorf("Expected InstallationId %v, got %v", model.InstallationId, result.InstallationId)
    }
    if result.ProviderRepositoryId != model.ProviderRepositoryId {
        t.Errorf("Expected ProviderRepositoryId %v, got %v", model.ProviderRepositoryId, result.ProviderRepositoryId)
    }
    if result.ProviderBranch != model.ProviderBranch {
        t.Errorf("Expected ProviderBranch %v, got %v", model.ProviderBranch, result.ProviderBranch)
    }
    if result.ProviderRootDirectory != model.ProviderRootDirectory {
        t.Errorf("Expected ProviderRootDirectory %v, got %v", model.ProviderRootDirectory, result.ProviderRootDirectory)
    }
    if result.ProviderSilentMode != model.ProviderSilentMode {
        t.Errorf("Expected ProviderSilentMode %v, got %v", model.ProviderSilentMode, result.ProviderSilentMode)
    }
    if result.BuildSpecification != model.BuildSpecification {
        t.Errorf("Expected BuildSpecification %v, got %v", model.BuildSpecification, result.BuildSpecification)
    }
    if result.RuntimeSpecification != model.RuntimeSpecification {
        t.Errorf("Expected RuntimeSpecification %v, got %v", model.RuntimeSpecification, result.RuntimeSpecification)
    }
    if result.BuildRuntime != model.BuildRuntime {
        t.Errorf("Expected BuildRuntime %v, got %v", model.BuildRuntime, result.BuildRuntime)
    }
    if result.Adapter != model.Adapter {
        t.Errorf("Expected Adapter %v, got %v", model.Adapter, result.Adapter)
    }
    if result.FallbackFile != model.FallbackFile {
        t.Errorf("Expected FallbackFile %v, got %v", model.FallbackFile, result.FallbackFile)
    }}
