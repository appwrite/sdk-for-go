package models

import (
    "encoding/json"
    "testing"
)

func TestWebhookModel(t *testing.T) {
    model := Webhook{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Webhook",        Url: "https://example.com/webhook",        Events: []string{"test"},        Security: true,        HttpUser: "username",        HttpPass: "password",        SignatureKey: "ad3d581ca230e2b7059c545e5a",        Enabled: true,        Logs: "Failed to connect to remote server.",        Attempts: 10,    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Webhook
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
    if result.Url != model.Url {
        t.Errorf("Expected Url %v, got %v", model.Url, result.Url)
    }
    if result.Security != model.Security {
        t.Errorf("Expected Security %v, got %v", model.Security, result.Security)
    }
    if result.HttpUser != model.HttpUser {
        t.Errorf("Expected HttpUser %v, got %v", model.HttpUser, result.HttpUser)
    }
    if result.HttpPass != model.HttpPass {
        t.Errorf("Expected HttpPass %v, got %v", model.HttpPass, result.HttpPass)
    }
    if result.SignatureKey != model.SignatureKey {
        t.Errorf("Expected SignatureKey %v, got %v", model.SignatureKey, result.SignatureKey)
    }
    if result.Enabled != model.Enabled {
        t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
    }
    if result.Logs != model.Logs {
        t.Errorf("Expected Logs %v, got %v", model.Logs, result.Logs)
    }
    if result.Attempts != model.Attempts {
        t.Errorf("Expected Attempts %v, got %v", model.Attempts, result.Attempts)
    }}
