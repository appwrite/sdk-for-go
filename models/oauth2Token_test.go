package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2TokenModel(t *testing.T) {
	model := Oauth2Token{AccessToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9...", TokenType: "Bearer", ExpiresIn: 3600, RefreshToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9...", Scope: "openid email profile"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2Token
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.AccessToken != model.AccessToken {
		t.Errorf("Expected AccessToken %v, got %v", model.AccessToken, result.AccessToken)
	}
	if result.TokenType != model.TokenType {
		t.Errorf("Expected TokenType %v, got %v", model.TokenType, result.TokenType)
	}
	if result.ExpiresIn != model.ExpiresIn {
		t.Errorf("Expected ExpiresIn %v, got %v", model.ExpiresIn, result.ExpiresIn)
	}
	if result.RefreshToken != model.RefreshToken {
		t.Errorf("Expected RefreshToken %v, got %v", model.RefreshToken, result.RefreshToken)
	}
	if result.Scope != model.Scope {
		t.Errorf("Expected Scope %v, got %v", model.Scope, result.Scope)
	}
}

func TestOauth2TokenModelDefaults(t *testing.T) {
	var fresh Oauth2Token
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.AccessToken)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing access_token to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"Bearer\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.TokenType)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing token_type to default to %v, got %v", want, got)
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
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.RefreshToken)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing refresh_token to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Scope)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing scope to default to %v, got %v", want, got)
		}
	}

	reused := Oauth2Token{AccessToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9...", TokenType: "Bearer", ExpiresIn: 3600, RefreshToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9...", Scope: "openid email profile"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
