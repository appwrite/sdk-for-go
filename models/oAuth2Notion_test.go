package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOAuth2NotionModel(t *testing.T) {
	model := OAuth2Notion{Id: "github", Enabled: true, OauthClientId: "341d8700-0000-0000-0000-000000446ee3", OauthClientSecret: "secret_dLUr4b000000000000000000000000000000lFHAa9"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result OAuth2Notion
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
	if result.OauthClientId != model.OauthClientId {
		t.Errorf("Expected OauthClientId %v, got %v", model.OauthClientId, result.OauthClientId)
	}
	if result.OauthClientSecret != model.OauthClientSecret {
		t.Errorf("Expected OauthClientSecret %v, got %v", model.OauthClientSecret, result.OauthClientSecret)
	}
}

func TestOAuth2NotionModelDefaults(t *testing.T) {
	var fresh OAuth2Notion
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
		encoded, err := json.Marshal(fresh.OauthClientId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing oauthClientId to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.OauthClientSecret)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing oauthClientSecret to default to %v, got %v", want, got)
		}
	}

	reused := OAuth2Notion{Id: "github", Enabled: true, OauthClientId: "341d8700-0000-0000-0000-000000446ee3", OauthClientSecret: "secret_dLUr4b000000000000000000000000000000lFHAa9"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
