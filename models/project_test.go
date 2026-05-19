package models

import (
    "encoding/json"
    "testing"
)

func TestProjectModel(t *testing.T) {
    model := Project{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "New Project",        TeamId: "1592981250",        DevKeys: []DevKey{DevKey{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Dev API Key",        Expire: "2020-10-15T06:38:00.000+00:00",        Secret: "919c2d18fb5d4...a2ae413da83346ad2",        AccessedAt: "2020-10-15T06:38:00.000+00:00",        Sdks: []string{"test"},    },
            },        SmtpEnabled: true,        SmtpSenderName: "John Appwrite",        SmtpSenderEmail: "john@appwrite.io",        SmtpReplyToName: "Support Team",        SmtpReplyToEmail: "support@appwrite.io",        SmtpHost: "mail.appwrite.io",        SmtpPort: 25,        SmtpUsername: "emailuser",        SmtpPassword: "string",        SmtpSecure: "tls",        PingCount: 1,        PingedAt: "2020-10-15T06:38:00.000+00:00",        Labels: []string{"test"},        Status: "active",        AuthMethods: []ProjectAuthMethod{ProjectAuthMethod{        Id: "email-password",        Enabled: true,    },
            },        Services: []ProjectService{ProjectService{        Id: "sites",        Enabled: true,    },
            },        Protocols: []ProjectProtocol{ProjectProtocol{        Id: "graphql",        Enabled: true,    },
            },        Region: "fra",        BillingLimits: BillingLimits{        Bandwidth: 5,        Storage: 150,        Users: 200000,        Executions: 750000,        GBHours: 100,        ImageTransformations: 100,        AuthPhone: 10,        BudgetLimit: 100,    },        Blocks: []Block{Block{        CreatedAt: "2020-10-15T06:38:00.000+00:00",        ResourceType: "project",        ResourceId: "5e5ea5c16897e",        ProjectName: "My Project",        Region: "fra",        OrganizationName: "Acme Inc.",        OrganizationId: "5e5ea5c16897e",        BillingPlan: "pro",    },
            },        ConsoleAccessedAt: "2020-10-15T06:38:00.000+00:00",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Project
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
    if result.TeamId != model.TeamId {
        t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
    }
    if result.SmtpEnabled != model.SmtpEnabled {
        t.Errorf("Expected SmtpEnabled %v, got %v", model.SmtpEnabled, result.SmtpEnabled)
    }
    if result.SmtpSenderName != model.SmtpSenderName {
        t.Errorf("Expected SmtpSenderName %v, got %v", model.SmtpSenderName, result.SmtpSenderName)
    }
    if result.SmtpSenderEmail != model.SmtpSenderEmail {
        t.Errorf("Expected SmtpSenderEmail %v, got %v", model.SmtpSenderEmail, result.SmtpSenderEmail)
    }
    if result.SmtpReplyToName != model.SmtpReplyToName {
        t.Errorf("Expected SmtpReplyToName %v, got %v", model.SmtpReplyToName, result.SmtpReplyToName)
    }
    if result.SmtpReplyToEmail != model.SmtpReplyToEmail {
        t.Errorf("Expected SmtpReplyToEmail %v, got %v", model.SmtpReplyToEmail, result.SmtpReplyToEmail)
    }
    if result.SmtpHost != model.SmtpHost {
        t.Errorf("Expected SmtpHost %v, got %v", model.SmtpHost, result.SmtpHost)
    }
    if result.SmtpPort != model.SmtpPort {
        t.Errorf("Expected SmtpPort %v, got %v", model.SmtpPort, result.SmtpPort)
    }
    if result.SmtpUsername != model.SmtpUsername {
        t.Errorf("Expected SmtpUsername %v, got %v", model.SmtpUsername, result.SmtpUsername)
    }
    if result.SmtpPassword != model.SmtpPassword {
        t.Errorf("Expected SmtpPassword %v, got %v", model.SmtpPassword, result.SmtpPassword)
    }
    if result.SmtpSecure != model.SmtpSecure {
        t.Errorf("Expected SmtpSecure %v, got %v", model.SmtpSecure, result.SmtpSecure)
    }
    if result.PingCount != model.PingCount {
        t.Errorf("Expected PingCount %v, got %v", model.PingCount, result.PingCount)
    }
    if result.PingedAt != model.PingedAt {
        t.Errorf("Expected PingedAt %v, got %v", model.PingedAt, result.PingedAt)
    }
    if result.Status != model.Status {
        t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
    }
    if result.Region != model.Region {
        t.Errorf("Expected Region %v, got %v", model.Region, result.Region)
    }
    if result.ConsoleAccessedAt != model.ConsoleAccessedAt {
        t.Errorf("Expected ConsoleAccessedAt %v, got %v", model.ConsoleAccessedAt, result.ConsoleAccessedAt)
    }}
