package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOauth2DeviceAuthorizationModel(t *testing.T) {
	model := Oauth2DeviceAuthorization{DeviceCode: "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a", UserCode: "ABCD-EFGH", VerificationUri: "https://appwrite.io/oauth2/device", VerificationUriComplete: "https://appwrite.io/oauth2/device?user_code=ABCD-EFGH", ExpiresIn: 900, Interval: 5}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Oauth2DeviceAuthorization
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.DeviceCode != model.DeviceCode {
		t.Errorf("Expected DeviceCode %v, got %v", model.DeviceCode, result.DeviceCode)
	}
	if result.UserCode != model.UserCode {
		t.Errorf("Expected UserCode %v, got %v", model.UserCode, result.UserCode)
	}
	if result.VerificationUri != model.VerificationUri {
		t.Errorf("Expected VerificationUri %v, got %v", model.VerificationUri, result.VerificationUri)
	}
	if result.VerificationUriComplete != model.VerificationUriComplete {
		t.Errorf("Expected VerificationUriComplete %v, got %v", model.VerificationUriComplete, result.VerificationUriComplete)
	}
	if result.ExpiresIn != model.ExpiresIn {
		t.Errorf("Expected ExpiresIn %v, got %v", model.ExpiresIn, result.ExpiresIn)
	}
	if result.Interval != model.Interval {
		t.Errorf("Expected Interval %v, got %v", model.Interval, result.Interval)
	}
}

func TestOauth2DeviceAuthorizationModelDefaults(t *testing.T) {
	var fresh Oauth2DeviceAuthorization
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DeviceCode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing device_code to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.UserCode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing user_code to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.VerificationUri)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing verification_uri to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.VerificationUriComplete)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing verification_uri_complete to default to %v, got %v", want, got)
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
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Interval)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing interval to default to %v, got %v", want, got)
		}
	}

	reused := Oauth2DeviceAuthorization{DeviceCode: "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a", UserCode: "ABCD-EFGH", VerificationUri: "https://appwrite.io/oauth2/device", VerificationUriComplete: "https://appwrite.io/oauth2/device?user_code=ABCD-EFGH", ExpiresIn: 900, Interval: 5}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
