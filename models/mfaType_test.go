package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMfaTypeModel(t *testing.T) {
	model := MfaType{Secret: "[SHARED_SECRET]", Uri: "otpauth://totp/appwrite:user@example.com?secret=[SHARED_SECRET]&issuer=appwrite"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result MfaType
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Secret != model.Secret {
		t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
	}
	if result.Uri != model.Uri {
		t.Errorf("Expected Uri %v, got %v", model.Uri, result.Uri)
	}
}

func TestMfaTypeModelDefaults(t *testing.T) {
	var fresh MfaType
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Secret)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing secret to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Uri)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing uri to default to %v, got %v", want, got)
		}
	}

	reused := MfaType{Secret: "[SHARED_SECRET]", Uri: "otpauth://totp/appwrite:user@example.com?secret=[SHARED_SECRET]&issuer=appwrite"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
