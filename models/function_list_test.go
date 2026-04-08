package models

import (
    "encoding/json"
    "testing"
)

func TestFunctionListModel(t *testing.T) {
    model := FunctionList{        Total: 5,        Functions: []Function{Function{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Execute: []string{"test"},        Name: "My Function",        Enabled: true,        Live: true,        Logging: true,        Runtime: "python-3.8",        DeploymentRetention: 7,        DeploymentId: "5e5ea5c16897e",        DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00",        LatestDeploymentId: "5e5ea5c16897e",        LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00",        LatestDeploymentStatus: "ready",        Scopes: []string{"test"},        Vars: []Variable{Variable{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "API_KEY",        Value: "myPa$$word1",        Secret: true,        ResourceType: "function",        ResourceId: "myAwesomeFunction",    },
            },        Events: []string{"test"},        Schedule: "5 4 * * *",        Timeout: 300,        Entrypoint: "index.js",        Commands: "npm install",        Version: "v2",        InstallationId: "6m40at4ejk5h2u9s1hboo",        ProviderRepositoryId: "appwrite",        ProviderBranch: "main",        ProviderRootDirectory: "functions/helloWorld",        ProviderSilentMode: true,        BuildSpecification: "s-1vcpu-512mb",        RuntimeSpecification: "s-1vcpu-512mb",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result FunctionList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
