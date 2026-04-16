package models

import (
    "encoding/json"
    "testing"
)

func TestSiteListModel(t *testing.T) {
    model := SiteList{        Total: 5,        Sites: []Site{Site{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Site",        Enabled: true,        Live: true,        Logging: true,        Framework: "react",        DeploymentRetention: 7,        DeploymentId: "5e5ea5c16897e",        DeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00",        DeploymentScreenshotLight: "5e5ea5c16897e",        DeploymentScreenshotDark: "5e5ea5c16897e",        LatestDeploymentId: "5e5ea5c16897e",        LatestDeploymentCreatedAt: "2020-10-15T06:38:00.000+00:00",        LatestDeploymentStatus: "ready",        Vars: []Variable{Variable{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Key: "API_KEY",        Value: "myPa$$word1",        Secret: true,        ResourceType: "function",        ResourceId: "myAwesomeFunction",    },
            },        Timeout: 300,        InstallCommand: "npm install",        BuildCommand: "npm run build",        StartCommand: "node custom-server.mjs",        OutputDirectory: "build",        InstallationId: "6m40at4ejk5h2u9s1hboo",        ProviderRepositoryId: "appwrite",        ProviderBranch: "main",        ProviderRootDirectory: "sites/helloWorld",        ProviderSilentMode: true,        BuildSpecification: "s-1vcpu-512mb",        RuntimeSpecification: "s-1vcpu-512mb",        BuildRuntime: "node-22",        Adapter: "static",        FallbackFile: "index.html",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result SiteList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
