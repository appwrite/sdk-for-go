package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAlertListModel(t *testing.T) {
	model := AlertList{Total: 5, Alerts: []Alert{Alert{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", MonitorId: "5e5ea5c16897e", Name: "Bandwidth spike", Status: "firing", AcknowledgedBy: "5e5ea5c16897e", AcknowledgedByName: "Jane Doe", EscalationId: "5e5ea5c16897e", EscalationStep: 1, EscalationRound: 0, Rules: []string{"test"}, Severity: "warning", Threshold: 5000000, Value: 7250000, Peak: 9100000, Type: "usage", ResourceType: "database", ResourceId: "5e5ea5c16897e", Notifications: 1}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AlertList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestAlertListModelDefaults(t *testing.T) {
	var fresh AlertList
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

	reused := AlertList{Total: 5, Alerts: []Alert{Alert{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", MonitorId: "5e5ea5c16897e", Name: "Bandwidth spike", Status: "firing", AcknowledgedBy: "5e5ea5c16897e", AcknowledgedByName: "Jane Doe", EscalationId: "5e5ea5c16897e", EscalationStep: 1, EscalationRound: 0, Rules: []string{"test"}, Severity: "warning", Threshold: 5000000, Value: 7250000, Peak: 9100000, Type: "usage", ResourceType: "database", ResourceId: "5e5ea5c16897e", Notifications: 1}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
