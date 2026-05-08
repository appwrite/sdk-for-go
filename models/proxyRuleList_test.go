package models

import (
    "encoding/json"
    "testing"
)

func TestProxyRuleListModel(t *testing.T) {
    model := ProxyRuleList{        Total: 5,        Rules: []ProxyRule{ProxyRule{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Domain: "appwrite.company.com",        Type: "deployment",        Trigger: "manual",        RedirectUrl: "https://appwrite.io/docs",        RedirectStatusCode: 301,        DeploymentId: "n3u9feiwmf",        DeploymentResourceId: "n3u9feiwmf",        DeploymentVcsProviderBranch: "main",        Status: "verified",        Logs: "Verification of DNS records failed with DNS resolver 8.8.8.8. Domain stage.myapp.com does not have DNS record.",        RenewAt: "datetime",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ProxyRuleList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
