package models

import (
    "encoding/json"
    "testing"
)

func TestAppsListModel(t *testing.T) {
    model := AppsList{        Total: 5,        Apps: []App{App{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Application",        Description: "Connect your workspace to My Application.",        ClientUri: "https://example.com",        LogoUri: "https://example.com/logo.png",        PrivacyPolicyUrl: "https://example.com/privacy",        TermsUrl: "https://example.com/terms",        Contacts: []string{"test"},        Tagline: "Automate your workspace.",        Tags: []string{"test"},        Labels: []string{"test"},        Images: []string{"test"},        SupportUrl: "https://example.com/support",        DataDeletionUrl: "https://example.com/data-deletion",        RedirectUris: []string{"test"},        PostLogoutRedirectUris: []string{"test"},        Enabled: true,        Type: "confidential",        DeviceFlow: true,        TeamId: "5e5ea5c16897e",        UserId: "5e5ea5c16897e",        InstallationScopes: []string{"test"},        InstallationRedirectUrl: "https://example.com/setup",        Secrets: []AppSecret{AppSecret{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        AppId: "5e5ea5c16897e",        Secret: "string",        Hint: "f5c6c7",        CreatedById: "5e5ea5c16897e",        CreatedByName: "Walter White",    },
            },    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result AppsList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
