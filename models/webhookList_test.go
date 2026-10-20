package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWebhookListModel(t *testing.T) {
	model := WebhookList{Total: 5, Webhooks: []Webhook{Webhook{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "My Webhook", Description: "Notifies the billing service about new users.", Url: "https://example.com/webhook", Events: []string{"test"}, Tls: true, AuthUsername: "username", AuthPassword: "webhook-password", Secret: "ad3d581ca230e2b7059c545e5a", Enabled: true, Logs: "Failed to connect to remote server.", Attempts: 10}}}

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
	}
}

func TestWebhookListModelDefaults(t *testing.T) {
	var fresh WebhookList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := WebhookList{Total: 5, Webhooks: []Webhook{Webhook{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "My Webhook", Description: "Notifies the billing service about new users.", Url: "https://example.com/webhook", Events: []string{"test"}, Tls: true, AuthUsername: "username", AuthPassword: "webhook-password", Secret: "ad3d581ca230e2b7059c545e5a", Enabled: true, Logs: "Failed to connect to remote server.", Attempts: 10}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
