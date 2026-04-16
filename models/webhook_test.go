package models

import (
    "encoding/json"
    "testing"
)

func TestWebhookModel(t *testing.T) {
    model := Webhook{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Webhook",        Url: "https://example.com/webhook",        Events: []string{"test"},        Tls: true,        AuthUsername: "username",        AuthPassword: "password",        Secret: "ad3d581ca230e2b7059c545e5a",        Enabled: true,        Logs: "Failed to connect to remote server.",        Attempts: 10,    }

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
    if result.Tls != model.Tls {
        t.Errorf("Expected Tls %v, got %v", model.Tls, result.Tls)
    }
    if result.AuthUsername != model.AuthUsername {
        t.Errorf("Expected AuthUsername %v, got %v", model.AuthUsername, result.AuthUsername)
    }
    if result.AuthPassword != model.AuthPassword {
        t.Errorf("Expected AuthPassword %v, got %v", model.AuthPassword, result.AuthPassword)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
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
