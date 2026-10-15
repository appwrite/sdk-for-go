package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2PARModel(t *testing.T) {
	model := Oauth2PAR{RequestUri: "urn:appwrite:oauth2:request:5e5ea5c16897e", ExpiresIn: 600}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2PAR
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.RequestUri != model.RequestUri {
		t.Errorf("Expected RequestUri %v, got %v", model.RequestUri, result.RequestUri)
	}
	if result.ExpiresIn != model.ExpiresIn {
		t.Errorf("Expected ExpiresIn %v, got %v", model.ExpiresIn, result.ExpiresIn)
	}
}

func TestOauth2PARModelDefaults(t *testing.T) {
	var fresh Oauth2PAR
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RequestUri)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing request_uri to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ExpiresIn)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing expires_in to default to %v, got %v", want, got)
		}
	}

	reused := Oauth2PAR{RequestUri: "urn:appwrite:oauth2:request:5e5ea5c16897e", ExpiresIn: 600}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
