package models

import (
	"encoding/json"
	"testing"
)

func TestOauth2DeviceAuthorizationModel(t *testing.T) {
	model := Oauth2DeviceAuthorization{DeviceCode: "5f3c8d2a1b9e4f7a6c8b2d1e9f4a7b3c5d8e1f2a9b4c7d6e3f5a8b1c4d7e2f9a", UserCode: "ABCD-EFGH", VerificationUri: "https://cloud.appwrite.io/console/oauth2/device", VerificationUriComplete: "https://cloud.appwrite.io/console/oauth2/device?user_code=ABCD-EFGH", ExpiresIn: 900, Interval: 5}

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
