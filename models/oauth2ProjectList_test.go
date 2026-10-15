package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2ProjectListModel(t *testing.T) {
	model := Oauth2ProjectList{Total: 5, Projects: []Oauth2Project{Oauth2Project{Id: "5e5ea5c16897e", Region: "fra", Endpoint: "https://fra.cloud.appwrite.io/v1"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2ProjectList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestOauth2ProjectListModelDefaults(t *testing.T) {
	var fresh Oauth2ProjectList
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

	reused := Oauth2ProjectList{Total: 5, Projects: []Oauth2Project{Oauth2Project{Id: "5e5ea5c16897e", Region: "fra", Endpoint: "https://fra.cloud.appwrite.io/v1"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
