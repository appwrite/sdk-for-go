package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOAuth2PaypalModel(t *testing.T) {
	model := OAuth2Paypal{Id: "github", Enabled: true, ClientId: "AdhIEG7-000000000000-0000000000000000000000000000000-0000000000000000000000-2pyB", SecretKey: "EH8KCXtew--000000000000000000000000000000000000000_C-1_5UP_000000000000000CB7KDp"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2Paypal
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
	if result.SecretKey != model.SecretKey {
		t.Errorf("Expected SecretKey %v, got %v", model.SecretKey, result.SecretKey)
	}
}

func TestOAuth2PaypalModelDefaults(t *testing.T) {
	var fresh OAuth2Paypal
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
		encoded, err := json.Marshal(fresh.SecretKey)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing secretKey to default to %v, got %v", want, got)
		}
	}

	reused := OAuth2Paypal{Id: "github", Enabled: true, ClientId: "AdhIEG7-000000000000-0000000000000000000000000000000-0000000000000000000000-2pyB", SecretKey: "EH8KCXtew--000000000000000000000000000000000000000_C-1_5UP_000000000000000CB7KDp"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
