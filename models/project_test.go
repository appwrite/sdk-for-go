package models

import (
    "encoding/json"
    "testing"
)

func TestProjectModel(t *testing.T) {
    model := Project{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "New Project",        Description: "This is a new project.",        TeamId: "1592981250",        Logo: "5f5c451b403cb",        Url: "5f5c451b403cb",        LegalName: "Company LTD.",        LegalCountry: "US",        LegalState: "New York",        LegalCity: "New York City.",        LegalAddress: "620 Eighth Avenue, New York, NY 10018",        LegalTaxId: "131102020",        AuthDuration: 60,        AuthLimit: 100,        AuthSessionsLimit: 10,        AuthPasswordHistory: 5,        AuthPasswordDictionary: true,        AuthPersonalDataCheck: true,        AuthDisposableEmails: true,        AuthCanonicalEmails: true,        AuthFreeEmails: true,        AuthMockNumbers: []MockNumber{MockNumber{        Phone: "+1612842323",        Otp: "123456",    },
            },        AuthSessionAlerts: true,        AuthMembershipsUserName: true,        AuthMembershipsUserEmail: true,        AuthMembershipsMfa: true,        AuthInvalidateSessions: true,        OAuthProviders: []AuthProvider{AuthProvider{        Key: "github",        Name: "GitHub",        AppId: "259125845563242502",        Secret: "Bpw_g9c2TGXxfgLshDbSaL8tsCcqgczQ",        Enabled: true,    },
            },        Platforms: []interface{}{},        Webhooks: []Webhook{Webhook{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Webhook",        Url: "https://example.com/webhook",        Events: []string{"test"},        Tls: true,        AuthUsername: "username",        AuthPassword: "password",        Secret: "ad3d581ca230e2b7059c545e5a",        Enabled: true,        Logs: "Failed to connect to remote server.",        Attempts: 10,    },
            },        Keys: []Key{Key{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My API Key",        Expire: "2020-10-15T06:38:00.000+00:00",        Scopes: []string{"test"},        Secret: "919c2d18fb5d4...a2ae413da83346ad2",        AccessedAt: "2020-10-15T06:38:00.000+00:00",        Sdks: []string{"test"},    },
            },        DevKeys: []DevKey{DevKey{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Dev API Key",        Expire: "2020-10-15T06:38:00.000+00:00",        Secret: "919c2d18fb5d4...a2ae413da83346ad2",        AccessedAt: "2020-10-15T06:38:00.000+00:00",        Sdks: []string{"test"},    },
            },        SmtpEnabled: true,        SmtpSenderName: "John Appwrite",        SmtpSenderEmail: "john@appwrite.io",        SmtpReplyTo: "support@appwrite.io",        SmtpHost: "mail.appwrite.io",        SmtpPort: 25,        SmtpUsername: "emailuser",        SmtpPassword: "securepassword",        SmtpSecure: "tls",        PingCount: 1,        PingedAt: "2020-10-15T06:38:00.000+00:00",        Labels: []string{"test"},        Status: "active",        AuthEmailPassword: true,        AuthUsersAuthMagicURL: true,        AuthEmailOtp: true,        AuthAnonymous: true,        AuthInvites: true,        AuthJWT: true,        AuthPhone: true,        ServiceStatusForAccount: true,        ServiceStatusForAvatars: true,        ServiceStatusForDatabases: true,        ServiceStatusForTablesdb: true,        ServiceStatusForLocale: true,        ServiceStatusForHealth: true,        ServiceStatusForProject: true,        ServiceStatusForStorage: true,        ServiceStatusForTeams: true,        ServiceStatusForUsers: true,        ServiceStatusForVcs: true,        ServiceStatusForSites: true,        ServiceStatusForFunctions: true,        ServiceStatusForProxy: true,        ServiceStatusForGraphql: true,        ServiceStatusForMigrations: true,        ServiceStatusForMessaging: true,        ProtocolStatusForRest: true,        ProtocolStatusForGraphql: true,        ProtocolStatusForWebsocket: true,        Region: "fra",        BillingLimits: BillingLimits{        Bandwidth: 5,        Storage: 150,        Users: 200000,        Executions: 750000,        GBHours: 100,        ImageTransformations: 100,        AuthPhone: 10,        BudgetLimit: 100,    },        Blocks: []Block{Block{        CreatedAt: "2020-10-15T06:38:00.000+00:00",        ResourceType: "project",        ResourceId: "5e5ea5c16897e",    },
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
    if result.Description != model.Description {
        t.Errorf("Expected Description %v, got %v", model.Description, result.Description)
    }
    if result.TeamId != model.TeamId {
        t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
    }
    if result.Logo != model.Logo {
        t.Errorf("Expected Logo %v, got %v", model.Logo, result.Logo)
    }
    if result.Url != model.Url {
        t.Errorf("Expected Url %v, got %v", model.Url, result.Url)
    }
    if result.LegalName != model.LegalName {
        t.Errorf("Expected LegalName %v, got %v", model.LegalName, result.LegalName)
    }
    if result.LegalCountry != model.LegalCountry {
        t.Errorf("Expected LegalCountry %v, got %v", model.LegalCountry, result.LegalCountry)
    }
    if result.LegalState != model.LegalState {
        t.Errorf("Expected LegalState %v, got %v", model.LegalState, result.LegalState)
    }
    if result.LegalCity != model.LegalCity {
        t.Errorf("Expected LegalCity %v, got %v", model.LegalCity, result.LegalCity)
    }
    if result.LegalAddress != model.LegalAddress {
        t.Errorf("Expected LegalAddress %v, got %v", model.LegalAddress, result.LegalAddress)
    }
    if result.LegalTaxId != model.LegalTaxId {
        t.Errorf("Expected LegalTaxId %v, got %v", model.LegalTaxId, result.LegalTaxId)
    }
    if result.AuthDuration != model.AuthDuration {
        t.Errorf("Expected AuthDuration %v, got %v", model.AuthDuration, result.AuthDuration)
    }
    if result.AuthLimit != model.AuthLimit {
        t.Errorf("Expected AuthLimit %v, got %v", model.AuthLimit, result.AuthLimit)
    }
    if result.AuthSessionsLimit != model.AuthSessionsLimit {
        t.Errorf("Expected AuthSessionsLimit %v, got %v", model.AuthSessionsLimit, result.AuthSessionsLimit)
    }
    if result.AuthPasswordHistory != model.AuthPasswordHistory {
        t.Errorf("Expected AuthPasswordHistory %v, got %v", model.AuthPasswordHistory, result.AuthPasswordHistory)
    }
    if result.AuthPasswordDictionary != model.AuthPasswordDictionary {
        t.Errorf("Expected AuthPasswordDictionary %v, got %v", model.AuthPasswordDictionary, result.AuthPasswordDictionary)
    }
    if result.AuthPersonalDataCheck != model.AuthPersonalDataCheck {
        t.Errorf("Expected AuthPersonalDataCheck %v, got %v", model.AuthPersonalDataCheck, result.AuthPersonalDataCheck)
    }
    if result.AuthDisposableEmails != model.AuthDisposableEmails {
        t.Errorf("Expected AuthDisposableEmails %v, got %v", model.AuthDisposableEmails, result.AuthDisposableEmails)
    }
    if result.AuthCanonicalEmails != model.AuthCanonicalEmails {
        t.Errorf("Expected AuthCanonicalEmails %v, got %v", model.AuthCanonicalEmails, result.AuthCanonicalEmails)
    }
    if result.AuthFreeEmails != model.AuthFreeEmails {
        t.Errorf("Expected AuthFreeEmails %v, got %v", model.AuthFreeEmails, result.AuthFreeEmails)
    }
    if result.AuthSessionAlerts != model.AuthSessionAlerts {
        t.Errorf("Expected AuthSessionAlerts %v, got %v", model.AuthSessionAlerts, result.AuthSessionAlerts)
    }
    if result.AuthMembershipsUserName != model.AuthMembershipsUserName {
        t.Errorf("Expected AuthMembershipsUserName %v, got %v", model.AuthMembershipsUserName, result.AuthMembershipsUserName)
    }
    if result.AuthMembershipsUserEmail != model.AuthMembershipsUserEmail {
        t.Errorf("Expected AuthMembershipsUserEmail %v, got %v", model.AuthMembershipsUserEmail, result.AuthMembershipsUserEmail)
    }
    if result.AuthMembershipsMfa != model.AuthMembershipsMfa {
        t.Errorf("Expected AuthMembershipsMfa %v, got %v", model.AuthMembershipsMfa, result.AuthMembershipsMfa)
    }
    if result.AuthInvalidateSessions != model.AuthInvalidateSessions {
        t.Errorf("Expected AuthInvalidateSessions %v, got %v", model.AuthInvalidateSessions, result.AuthInvalidateSessions)
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
    if result.SmtpReplyTo != model.SmtpReplyTo {
        t.Errorf("Expected SmtpReplyTo %v, got %v", model.SmtpReplyTo, result.SmtpReplyTo)
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
    if result.AuthEmailPassword != model.AuthEmailPassword {
        t.Errorf("Expected AuthEmailPassword %v, got %v", model.AuthEmailPassword, result.AuthEmailPassword)
    }
    if result.AuthUsersAuthMagicURL != model.AuthUsersAuthMagicURL {
        t.Errorf("Expected AuthUsersAuthMagicURL %v, got %v", model.AuthUsersAuthMagicURL, result.AuthUsersAuthMagicURL)
    }
    if result.AuthEmailOtp != model.AuthEmailOtp {
        t.Errorf("Expected AuthEmailOtp %v, got %v", model.AuthEmailOtp, result.AuthEmailOtp)
    }
    if result.AuthAnonymous != model.AuthAnonymous {
        t.Errorf("Expected AuthAnonymous %v, got %v", model.AuthAnonymous, result.AuthAnonymous)
    }
    if result.AuthInvites != model.AuthInvites {
        t.Errorf("Expected AuthInvites %v, got %v", model.AuthInvites, result.AuthInvites)
    }
    if result.AuthJWT != model.AuthJWT {
        t.Errorf("Expected AuthJWT %v, got %v", model.AuthJWT, result.AuthJWT)
    }
    if result.AuthPhone != model.AuthPhone {
        t.Errorf("Expected AuthPhone %v, got %v", model.AuthPhone, result.AuthPhone)
    }
    if result.ServiceStatusForAccount != model.ServiceStatusForAccount {
        t.Errorf("Expected ServiceStatusForAccount %v, got %v", model.ServiceStatusForAccount, result.ServiceStatusForAccount)
    }
    if result.ServiceStatusForAvatars != model.ServiceStatusForAvatars {
        t.Errorf("Expected ServiceStatusForAvatars %v, got %v", model.ServiceStatusForAvatars, result.ServiceStatusForAvatars)
    }
    if result.ServiceStatusForDatabases != model.ServiceStatusForDatabases {
        t.Errorf("Expected ServiceStatusForDatabases %v, got %v", model.ServiceStatusForDatabases, result.ServiceStatusForDatabases)
    }
    if result.ServiceStatusForTablesdb != model.ServiceStatusForTablesdb {
        t.Errorf("Expected ServiceStatusForTablesdb %v, got %v", model.ServiceStatusForTablesdb, result.ServiceStatusForTablesdb)
    }
    if result.ServiceStatusForLocale != model.ServiceStatusForLocale {
        t.Errorf("Expected ServiceStatusForLocale %v, got %v", model.ServiceStatusForLocale, result.ServiceStatusForLocale)
    }
    if result.ServiceStatusForHealth != model.ServiceStatusForHealth {
        t.Errorf("Expected ServiceStatusForHealth %v, got %v", model.ServiceStatusForHealth, result.ServiceStatusForHealth)
    }
    if result.ServiceStatusForProject != model.ServiceStatusForProject {
        t.Errorf("Expected ServiceStatusForProject %v, got %v", model.ServiceStatusForProject, result.ServiceStatusForProject)
    }
    if result.ServiceStatusForStorage != model.ServiceStatusForStorage {
        t.Errorf("Expected ServiceStatusForStorage %v, got %v", model.ServiceStatusForStorage, result.ServiceStatusForStorage)
    }
    if result.ServiceStatusForTeams != model.ServiceStatusForTeams {
        t.Errorf("Expected ServiceStatusForTeams %v, got %v", model.ServiceStatusForTeams, result.ServiceStatusForTeams)
    }
    if result.ServiceStatusForUsers != model.ServiceStatusForUsers {
        t.Errorf("Expected ServiceStatusForUsers %v, got %v", model.ServiceStatusForUsers, result.ServiceStatusForUsers)
    }
    if result.ServiceStatusForVcs != model.ServiceStatusForVcs {
        t.Errorf("Expected ServiceStatusForVcs %v, got %v", model.ServiceStatusForVcs, result.ServiceStatusForVcs)
    }
    if result.ServiceStatusForSites != model.ServiceStatusForSites {
        t.Errorf("Expected ServiceStatusForSites %v, got %v", model.ServiceStatusForSites, result.ServiceStatusForSites)
    }
    if result.ServiceStatusForFunctions != model.ServiceStatusForFunctions {
        t.Errorf("Expected ServiceStatusForFunctions %v, got %v", model.ServiceStatusForFunctions, result.ServiceStatusForFunctions)
    }
    if result.ServiceStatusForProxy != model.ServiceStatusForProxy {
        t.Errorf("Expected ServiceStatusForProxy %v, got %v", model.ServiceStatusForProxy, result.ServiceStatusForProxy)
    }
    if result.ServiceStatusForGraphql != model.ServiceStatusForGraphql {
        t.Errorf("Expected ServiceStatusForGraphql %v, got %v", model.ServiceStatusForGraphql, result.ServiceStatusForGraphql)
    }
    if result.ServiceStatusForMigrations != model.ServiceStatusForMigrations {
        t.Errorf("Expected ServiceStatusForMigrations %v, got %v", model.ServiceStatusForMigrations, result.ServiceStatusForMigrations)
    }
    if result.ServiceStatusForMessaging != model.ServiceStatusForMessaging {
        t.Errorf("Expected ServiceStatusForMessaging %v, got %v", model.ServiceStatusForMessaging, result.ServiceStatusForMessaging)
    }
    if result.ProtocolStatusForRest != model.ProtocolStatusForRest {
        t.Errorf("Expected ProtocolStatusForRest %v, got %v", model.ProtocolStatusForRest, result.ProtocolStatusForRest)
    }
    if result.ProtocolStatusForGraphql != model.ProtocolStatusForGraphql {
        t.Errorf("Expected ProtocolStatusForGraphql %v, got %v", model.ProtocolStatusForGraphql, result.ProtocolStatusForGraphql)
    }
    if result.ProtocolStatusForWebsocket != model.ProtocolStatusForWebsocket {
        t.Errorf("Expected ProtocolStatusForWebsocket %v, got %v", model.ProtocolStatusForWebsocket, result.ProtocolStatusForWebsocket)
    }
    if result.Region != model.Region {
        t.Errorf("Expected Region %v, got %v", model.Region, result.Region)
    }
    if result.ConsoleAccessedAt != model.ConsoleAccessedAt {
        t.Errorf("Expected ConsoleAccessedAt %v, got %v", model.ConsoleAccessedAt, result.ConsoleAccessedAt)
    }}
