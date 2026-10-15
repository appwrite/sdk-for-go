package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAppInstallationListModel(t *testing.T) {
	model := AppInstallationList{Total: 5, Installations: []AppInstallation{AppInstallation{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", AppId: "5e5ea5c16897e", TeamId: "5e5ea5c16897e", Scopes: []string{"test"}, AuthorizationDetails: []interface{}{}, CreatedById: "5e5ea5c16897e", CreatedByName: "Walter White"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AppInstallationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestAppInstallationListModelDefaults(t *testing.T) {
	var fresh AppInstallationList
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

	reused := AppInstallationList{Total: 5, Installations: []AppInstallation{AppInstallation{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", AppId: "5e5ea5c16897e", TeamId: "5e5ea5c16897e", Scopes: []string{"test"}, AuthorizationDetails: []interface{}{}, CreatedById: "5e5ea5c16897e", CreatedByName: "Walter White"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
