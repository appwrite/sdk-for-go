package models

import (
    "encoding/json"
    "testing"
)

func TestDeploymentModel(t *testing.T) {
    model := Deployment{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Type: "vcs",        ResourceId: "5e5ea6g16897e",        ResourceType: "functions",        Entrypoint: "index.js",        SourceSize: 128,        BuildSize: 128,        TotalSize: 128,        BuildId: "5e5ea5c16897e",        Activate: true,        ScreenshotLight: "5e5ea5c16897e",        ScreenshotDark: "5e5ea5c16897e",        Status: "ready",        BuildLogs: "Compiling source files...",        BuildDuration: 128,        ProviderRepositoryName: "database",        ProviderRepositoryOwner: "utopia",        ProviderRepositoryUrl: "https://github.com/vermakhushboo/g4-node-function",        ProviderCommitHash: "7c3f25d",        ProviderCommitAuthorUrl: "https://github.com/vermakhushboo",        ProviderCommitAuthor: "Khushboo Verma",        ProviderCommitMessage: "Update index.js",        ProviderCommitUrl: "https://github.com/vermakhushboo/g4-node-function/commit/60c0416257a9cbcdd96b2d370c38d8f8d150ccfb",        ProviderBranch: "0.7.x",        ProviderBranchUrl: "https://github.com/vermakhushboo/appwrite/tree/0.7.x",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Deployment
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
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.ResourceId != model.ResourceId {
        t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
    }
    if result.ResourceType != model.ResourceType {
        t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
    }
    if result.Entrypoint != model.Entrypoint {
        t.Errorf("Expected Entrypoint %v, got %v", model.Entrypoint, result.Entrypoint)
    }
    if result.SourceSize != model.SourceSize {
        t.Errorf("Expected SourceSize %v, got %v", model.SourceSize, result.SourceSize)
    }
    if result.BuildSize != model.BuildSize {
        t.Errorf("Expected BuildSize %v, got %v", model.BuildSize, result.BuildSize)
    }
    if result.TotalSize != model.TotalSize {
        t.Errorf("Expected TotalSize %v, got %v", model.TotalSize, result.TotalSize)
    }
    if result.BuildId != model.BuildId {
        t.Errorf("Expected BuildId %v, got %v", model.BuildId, result.BuildId)
    }
    if result.Activate != model.Activate {
        t.Errorf("Expected Activate %v, got %v", model.Activate, result.Activate)
    }
    if result.ScreenshotLight != model.ScreenshotLight {
        t.Errorf("Expected ScreenshotLight %v, got %v", model.ScreenshotLight, result.ScreenshotLight)
    }
    if result.ScreenshotDark != model.ScreenshotDark {
        t.Errorf("Expected ScreenshotDark %v, got %v", model.ScreenshotDark, result.ScreenshotDark)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.BuildLogs != model.BuildLogs {
        t.Errorf("Expected BuildLogs %v, got %v", model.BuildLogs, result.BuildLogs)
    }
    if result.BuildDuration != model.BuildDuration {
        t.Errorf("Expected BuildDuration %v, got %v", model.BuildDuration, result.BuildDuration)
    }
    if result.ProviderRepositoryName != model.ProviderRepositoryName {
        t.Errorf("Expected ProviderRepositoryName %v, got %v", model.ProviderRepositoryName, result.ProviderRepositoryName)
    }
    if result.ProviderRepositoryOwner != model.ProviderRepositoryOwner {
        t.Errorf("Expected ProviderRepositoryOwner %v, got %v", model.ProviderRepositoryOwner, result.ProviderRepositoryOwner)
    }
    if result.ProviderRepositoryUrl != model.ProviderRepositoryUrl {
        t.Errorf("Expected ProviderRepositoryUrl %v, got %v", model.ProviderRepositoryUrl, result.ProviderRepositoryUrl)
    }
    if result.ProviderCommitHash != model.ProviderCommitHash {
        t.Errorf("Expected ProviderCommitHash %v, got %v", model.ProviderCommitHash, result.ProviderCommitHash)
    }
    if result.ProviderCommitAuthorUrl != model.ProviderCommitAuthorUrl {
        t.Errorf("Expected ProviderCommitAuthorUrl %v, got %v", model.ProviderCommitAuthorUrl, result.ProviderCommitAuthorUrl)
    }
    if result.ProviderCommitAuthor != model.ProviderCommitAuthor {
        t.Errorf("Expected ProviderCommitAuthor %v, got %v", model.ProviderCommitAuthor, result.ProviderCommitAuthor)
    }
    if result.ProviderCommitMessage != model.ProviderCommitMessage {
        t.Errorf("Expected ProviderCommitMessage %v, got %v", model.ProviderCommitMessage, result.ProviderCommitMessage)
    }
    if result.ProviderCommitUrl != model.ProviderCommitUrl {
        t.Errorf("Expected ProviderCommitUrl %v, got %v", model.ProviderCommitUrl, result.ProviderCommitUrl)
    }
    if result.ProviderBranch != model.ProviderBranch {
        t.Errorf("Expected ProviderBranch %v, got %v", model.ProviderBranch, result.ProviderBranch)
    }
    if result.ProviderBranchUrl != model.ProviderBranchUrl {
        t.Errorf("Expected ProviderBranchUrl %v, got %v", model.ProviderBranchUrl, result.ProviderBranchUrl)
    }}
