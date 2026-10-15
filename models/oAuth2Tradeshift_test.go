package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOAuth2TradeshiftModel(t *testing.T) {
	model := OAuth2Tradeshift{Id: "github", Enabled: true, Oauth2ClientId: "appwrite-test-org.appwrite-test-app", Oauth2ClientSecret: "7cb52700-0000-0000-0000-000000ca5b83"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2Tradeshift
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
	if result.Oauth2ClientId != model.Oauth2ClientId {
		t.Errorf("Expected Oauth2ClientId %v, got %v", model.Oauth2ClientId, result.Oauth2ClientId)
	}
	if result.Oauth2ClientSecret != model.Oauth2ClientSecret {
		t.Errorf("Expected Oauth2ClientSecret %v, got %v", model.Oauth2ClientSecret, result.Oauth2ClientSecret)
	}
}

func TestOAuth2TradeshiftModelDefaults(t *testing.T) {
	var fresh OAuth2Tradeshift
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
		encoded, err := json.Marshal(fresh.Oauth2ClientId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing oauth2ClientId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Oauth2ClientSecret)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing oauth2ClientSecret to default to %v, got %v", want, got)
		}
	}

	reused := OAuth2Tradeshift{Id: "github", Enabled: true, Oauth2ClientId: "appwrite-test-org.appwrite-test-app", Oauth2ClientSecret: "7cb52700-0000-0000-0000-000000ca5b83"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
