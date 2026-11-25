package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMonitorListModel(t *testing.T) {
	model := MonitorList{Total: 5, Monitors: []Monitor{Monitor{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "Bandwidth spike", ProjectId: "5e5ea5c16897e", Rules: []string{"test"}, Type: "usage", ResourceType: "database", ResourceId: "main", Window: 3600, Cooldown: 3600, Enabled: true, Notify: []string{"test"}, EscalationId: "5e5ea5c16897e", Severity: "warning", Status: "ok", AlertId: "5e5ea5c16897e", Interval: 60, Timeout: 10, Url: "https://example.com/health", Body: "string", Domain: "example.com", Resolver: "1.1.1.1", Host: "mail.example.com", Port: 587}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result MonitorList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestMonitorListModelDefaults(t *testing.T) {
	var fresh MonitorList
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

	reused := MonitorList{Total: 5, Monitors: []Monitor{Monitor{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "Bandwidth spike", ProjectId: "5e5ea5c16897e", Rules: []string{"test"}, Type: "usage", ResourceType: "database", ResourceId: "main", Window: 3600, Cooldown: 3600, Enabled: true, Notify: []string{"test"}, EscalationId: "5e5ea5c16897e", Severity: "warning", Status: "ok", AlertId: "5e5ea5c16897e", Interval: 60, Timeout: 10, Url: "https://example.com/health", Body: "string", Domain: "example.com", Resolver: "1.1.1.1", Host: "mail.example.com", Port: 587}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
