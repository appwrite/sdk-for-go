package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2OrganizationListModel(t *testing.T) {
	model := Oauth2OrganizationList{Total: 5, Organizations: []Oauth2Organization{Oauth2Organization{Id: "5e5ea5c16897e"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2OrganizationList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestOauth2OrganizationListModelDefaults(t *testing.T) {
	var fresh Oauth2OrganizationList
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

	reused := Oauth2OrganizationList{Total: 5, Organizations: []Oauth2Organization{Oauth2Organization{Id: "5e5ea5c16897e"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
