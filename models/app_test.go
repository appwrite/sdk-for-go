package models

import (
    "encoding/json"
    "testing"
)

func TestAppModel(t *testing.T) {
    model := App{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Application",        Description: "Connect your workspace to My Application.",        ClientUri: "https://example.com",        LogoUri: "https://example.com/logo.png",        PrivacyPolicyUrl: "https://example.com/privacy",        TermsUrl: "https://example.com/terms",        Contacts: []string{"test"},        Tagline: "Automate your workspace.",        Tags: []string{"test"},        Labels: []string{"test"},        Images: []string{"test"},        SupportUrl: "https://example.com/support",        DataDeletionUrl: "https://example.com/data-deletion",        RedirectUris: []string{"test"},        PostLogoutRedirectUris: []string{"test"},        Enabled: true,        Type: "confidential",        DeviceFlow: true,        TeamId: "5e5ea5c16897e",        UserId: "5e5ea5c16897e",        InstallationScopes: []string{"test"},        InstallationRedirectUrl: "https://example.com/setup",        Secrets: []AppSecret{AppSecret{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        AppId: "5e5ea5c16897e",        Secret: "string",        Hint: "f5c6c7",        CreatedById: "5e5ea5c16897e",        CreatedByName: "Walter White",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result App
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
    if result.Description != model.Description {
        t.Errorf("Expected Description %v, got %v", model.Description, result.Description)
    }
    if result.ClientUri != model.ClientUri {
        t.Errorf("Expected ClientUri %v, got %v", model.ClientUri, result.ClientUri)
    }
    if result.LogoUri != model.LogoUri {
        t.Errorf("Expected LogoUri %v, got %v", model.LogoUri, result.LogoUri)
    }
    if result.PrivacyPolicyUrl != model.PrivacyPolicyUrl {
        t.Errorf("Expected PrivacyPolicyUrl %v, got %v", model.PrivacyPolicyUrl, result.PrivacyPolicyUrl)
    }
    if result.TermsUrl != model.TermsUrl {
        t.Errorf("Expected TermsUrl %v, got %v", model.TermsUrl, result.TermsUrl)
    }
    if result.Tagline != model.Tagline {
        t.Errorf("Expected Tagline %v, got %v", model.Tagline, result.Tagline)
    }
    if result.SupportUrl != model.SupportUrl {
        t.Errorf("Expected SupportUrl %v, got %v", model.SupportUrl, result.SupportUrl)
    }
    if result.DataDeletionUrl != model.DataDeletionUrl {
        t.Errorf("Expected DataDeletionUrl %v, got %v", model.DataDeletionUrl, result.DataDeletionUrl)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }
    if result.Type != model.Type {
        t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
    }
    if result.DeviceFlow != model.DeviceFlow {
        t.Errorf("Expected DeviceFlow %v, got %v", model.DeviceFlow, result.DeviceFlow)
    }
    if result.TeamId != model.TeamId {
        t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
    }
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.InstallationRedirectUrl != model.InstallationRedirectUrl {
        t.Errorf("Expected InstallationRedirectUrl %v, got %v", model.InstallationRedirectUrl, result.InstallationRedirectUrl)
    }}
