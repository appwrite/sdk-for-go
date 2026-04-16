package models

import (
    "encoding/json"
    "testing"
)

func TestWebhookListModel(t *testing.T) {
    model := WebhookList{        Total: 5,        Webhooks: []Webhook{Webhook{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "My Webhook",        Url: "https://example.com/webhook",        Events: []string{"test"},        Tls: true,        AuthUsername: "username",        AuthPassword: "password",        Secret: "ad3d581ca230e2b7059c545e5a",        Enabled: true,        Logs: "Failed to connect to remote server.",        Attempts: 10,    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result WebhookList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
