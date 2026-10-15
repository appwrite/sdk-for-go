package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestActivityEventListModel(t *testing.T) {
	model := ActivityEventList{Total: 5, Events: []ActivityEvent{ActivityEvent{Id: "5e5ea5c16897e", ActorType: "user", ActorId: "610fc2f985ee0", ActorEmail: "john@appwrite.io", ActorName: "John Doe", ResourceParent: "database/ID", ResourceType: "collection", ResourceId: "610fc2f985ee0", Resource: "collections/610fc2f985ee0", Event: "account.sessions.create", UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36", Ip: "127.0.0.1", Mode: "admin", Country: "US", ContinentCode: "NA", City: "Mountain View", Subdivisions: "California", Isp: "Google", AutonomousSystemNumber: "15169", AutonomousSystemOrganization: "GOOGLE", ConnectionType: "cable", ConnectionUsageType: "residential", ConnectionOrganization: "Google LLC", Time: "2020-10-15T06:38:00.000+00:00", ProjectId: "610fc2f985ee0", TeamId: "610fc2f985ee0", Hostname: "appwrite.io", Sdk: "web", SdkVersion: "14.0.0"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result ActivityEventList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestActivityEventListModelDefaults(t *testing.T) {
	var fresh ActivityEventList
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

	reused := ActivityEventList{Total: 5, Events: []ActivityEvent{ActivityEvent{Id: "5e5ea5c16897e", ActorType: "user", ActorId: "610fc2f985ee0", ActorEmail: "john@appwrite.io", ActorName: "John Doe", ResourceParent: "database/ID", ResourceType: "collection", ResourceId: "610fc2f985ee0", Resource: "collections/610fc2f985ee0", Event: "account.sessions.create", UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36", Ip: "127.0.0.1", Mode: "admin", Country: "US", ContinentCode: "NA", City: "Mountain View", Subdivisions: "California", Isp: "Google", AutonomousSystemNumber: "15169", AutonomousSystemOrganization: "GOOGLE", ConnectionType: "cable", ConnectionUsageType: "residential", ConnectionOrganization: "Google LLC", Time: "2020-10-15T06:38:00.000+00:00", ProjectId: "610fc2f985ee0", TeamId: "610fc2f985ee0", Hostname: "appwrite.io", Sdk: "web", SdkVersion: "14.0.0"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
