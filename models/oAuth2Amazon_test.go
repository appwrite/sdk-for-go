package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOAuth2AmazonModel(t *testing.T) {
	model := OAuth2Amazon{Id: "github", Enabled: true, ClientId: "amzn1.application-oa2-client.87400c00000000000000000000063d5b2", ClientSecret: "79ffe4000000000000000000000000000000000000000000000000000002de55"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2Amazon
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
	if result.ClientId != model.ClientId {
		t.Errorf("Expected ClientId %v, got %v", model.ClientId, result.ClientId)
	}
	if result.ClientSecret != model.ClientSecret {
		t.Errorf("Expected ClientSecret %v, got %v", model.ClientSecret, result.ClientSecret)
	}
}

func TestOAuth2AmazonModelDefaults(t *testing.T) {
	var fresh OAuth2Amazon
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Id)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $id to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Enabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing enabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ClientId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing clientId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ClientSecret)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing clientSecret to default to %v, got %v", want, got)
		}
	}

	reused := OAuth2Amazon{Id: "github", Enabled: true, ClientId: "amzn1.application-oa2-client.87400c00000000000000000000063d5b2", ClientSecret: "79ffe4000000000000000000000000000000000000000000000000000002de55"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
