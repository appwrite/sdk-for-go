package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDeviceModel(t *testing.T) {
	model := Device{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Expire: "2020-10-15T06:38:00.000+00:00", ActiveSessionId: "5e5bb8c16897e", Sessions: []Session{Session{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5bb8c16897e", UserName: "John Doe", Expire: "2020-10-15T06:38:00.000+00:00", Provider: "email", ProviderUid: "user@example.com", ProviderAccessToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", ProviderAccessTokenExpiry: "2020-10-15T06:38:00.000+00:00", ProviderRefreshToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", Ip: "127.0.0.1", OsCode: "Mac", OsName: "Mac", OsVersion: "Mac", ClientType: "browser", ClientCode: "CM", ClientName: "Chrome Mobile iOS", ClientVersion: "84.0", ClientEngine: "WebKit", ClientEngineVersion: "605.1.15", DeviceName: "smartphone", DeviceBrand: "Google", DeviceModel: "Nexus 5", CountryCode: "US", CountryName: "United States", Current: true, Factors: []string{"test"}, Secret: "5e5bb8c16897e", MfaUpdatedAt: "2020-10-15T06:38:00.000+00:00"}}, Secret: "5e5bb8c16897e"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Device
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
	if result.Expire != model.Expire {
		t.Errorf("Expected Expire %v, got %v", model.Expire, result.Expire)
	}
	if result.ActiveSessionId != model.ActiveSessionId {
		t.Errorf("Expected ActiveSessionId %v, got %v", model.ActiveSessionId, result.ActiveSessionId)
	}
	if result.Secret != model.Secret {
		t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
	}
}

func TestDeviceModelDefaults(t *testing.T) {
	var fresh Device
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
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CreatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $createdAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.UpdatedAt)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing $updatedAt to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Expire)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing expire to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ActiveSessionId)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing activeSessionId to default to %v, got %v", want, got)
		}
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

	reused := Device{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Expire: "2020-10-15T06:38:00.000+00:00", ActiveSessionId: "5e5bb8c16897e", Sessions: []Session{Session{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5bb8c16897e", UserName: "John Doe", Expire: "2020-10-15T06:38:00.000+00:00", Provider: "email", ProviderUid: "user@example.com", ProviderAccessToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", ProviderAccessTokenExpiry: "2020-10-15T06:38:00.000+00:00", ProviderRefreshToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", Ip: "127.0.0.1", OsCode: "Mac", OsName: "Mac", OsVersion: "Mac", ClientType: "browser", ClientCode: "CM", ClientName: "Chrome Mobile iOS", ClientVersion: "84.0", ClientEngine: "WebKit", ClientEngineVersion: "605.1.15", DeviceName: "smartphone", DeviceBrand: "Google", DeviceModel: "Nexus 5", CountryCode: "US", CountryName: "United States", Current: true, Factors: []string{"test"}, Secret: "5e5bb8c16897e", MfaUpdatedAt: "2020-10-15T06:38:00.000+00:00"}}, Secret: "5e5bb8c16897e"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
