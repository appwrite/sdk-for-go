package models

import (
    "encoding/json"
    "testing"
)

func TestProxyRuleModel(t *testing.T) {
    model := ProxyRule{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Domain: "appwrite.company.com",        Type: "deployment",        Trigger: "manual",        RedirectUrl: "https://appwrite.io/docs",        RedirectStatusCode: 301,        DeploymentId: "n3u9feiwmf",        DeploymentResourceId: "n3u9feiwmf",        DeploymentVcsProviderBranch: "main",        Status: "verified",        Logs: "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",        RenewAt: "datetime",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ProxyRule
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
    if result.Domain != model.Domain {
        t.Errorf("Expected Domain %v, got %v", model.Domain, result.Domain)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.Trigger != model.Trigger {
        t.Errorf("Expected Trigger %v, got %v", model.Trigger, result.Trigger)
    }
    if result.RedirectUrl != model.RedirectUrl {
        t.Errorf("Expected RedirectUrl %v, got %v", model.RedirectUrl, result.RedirectUrl)
    }
    if result.RedirectStatusCode != model.RedirectStatusCode {
        t.Errorf("Expected RedirectStatusCode %v, got %v", model.RedirectStatusCode, result.RedirectStatusCode)
    }
    if result.DeploymentId != model.DeploymentId {
        t.Errorf("Expected DeploymentId %v, got %v", model.DeploymentId, result.DeploymentId)
    }
    if result.DeploymentResourceId != model.DeploymentResourceId {
        t.Errorf("Expected DeploymentResourceId %v, got %v", model.DeploymentResourceId, result.DeploymentResourceId)
    }
    if result.DeploymentVcsProviderBranch != model.DeploymentVcsProviderBranch {
        t.Errorf("Expected DeploymentVcsProviderBranch %v, got %v", model.DeploymentVcsProviderBranch, result.DeploymentVcsProviderBranch)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.Logs != model.Logs {
        t.Errorf("Expected Logs %v, got %v", model.Logs, result.Logs)
    }
    if result.RenewAt != model.RenewAt {
        t.Errorf("Expected RenewAt %v, got %v", model.RenewAt, result.RenewAt)
    }}
