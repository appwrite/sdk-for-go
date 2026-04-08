package models

import (
    "encoding/json"
    "testing"
)

func TestFunctionModel(t *testing.T) {
    model := Function{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Execute: []string{"test"},        Name: "My Function",        Enabled: true,        Live: true,        Logging: true,        Runtime: "python-3.8",        DeploymentRetention: 7,        DeploymentId: "5e5ea5c16897e",        DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00",        LatestDeploymentId: "5e5ea5c16897e",        LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00",        LatestDeploymentStatus: "ready",        Scopes: []string{"test"},        Vars: []Variable{Variable{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "API_KEY",        Value: "myPa$$word1",        Secret: true,        ResourceType: "function",        ResourceId: "myAwesomeFunction",    },
            },        Events: []string{"test"},        Schedule: "5 4 * * *",        Timeout: 300,        Entrypoint: "index.js",        Commands: "npm install",        Version: "v2",        InstallationId: "6m40at4ejk5h2u9s1hboo",        ProviderRepositoryId: "appwrite",        ProviderBranch: "main",        ProviderRootDirectory: "functions/helloWorld",        ProviderSilentMode: true,        BuildSpecification: "s-1vcpu-512mb",        RuntimeSpecification: "s-1vcpu-512mb",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Function
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
    if result.Runtime != model.Runtime {
        t.Errorf("Expected Runtime %v, got %v", model.Runtime, result.Runtime)
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
    if result.LatestDeploymentId != model.LatestDeploymentId {
        t.Errorf("Expected LatestDeploymentId %v, got %v", model.LatestDeploymentId, result.LatestDeploymentId)
    }
    if result.LatestDeploymentCreatedAt != model.LatestDeploymentCreatedAt {
        t.Errorf("Expected LatestDeploymentCreatedAt %v, got %v", model.LatestDeploymentCreatedAt, result.LatestDeploymentCreatedAt)
    }
    if result.LatestDeploymentStatus != model.LatestDeploymentStatus {
        t.Errorf("Expected LatestDeploymentStatus %v, got %v", model.LatestDeploymentStatus, result.LatestDeploymentStatus)
    }
    if result.Schedule != model.Schedule {
        t.Errorf("Expected Schedule %v, got %v", model.Schedule, result.Schedule)
    }
    if result.Timeout != model.Timeout {
        t.Errorf("Expected Timeout %v, got %v", model.Timeout, result.Timeout)
    }
    if result.Entrypoint != model.Entrypoint {
        t.Errorf("Expected Entrypoint %v, got %v", model.Entrypoint, result.Entrypoint)
    }
    if result.Commands != model.Commands {
        t.Errorf("Expected Commands %v, got %v", model.Commands, result.Commands)
    }
    if result.Version != model.Version {
        t.Errorf("Expected Version %v, got %v", model.Version, result.Version)
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
    }}
