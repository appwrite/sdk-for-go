package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSessionListModel(t *testing.T) {
	model := SessionList{Total: 5, Sessions: []Session{Session{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5bb8c16897e", Expire: "2020-10-15T06:38:00.000+00:00", Provider: "email", ProviderUid: "user@example.com", ProviderAccessToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", ProviderAccessTokenExpiry: "2020-10-15T06:38:00.000+00:00", ProviderRefreshToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", Ip: "127.0.0.1", OsCode: "Mac", OsName: "Mac", OsVersion: "Mac", ClientType: "browser", ClientCode: "CM", ClientName: "Chrome Mobile iOS", ClientVersion: "84.0", ClientEngine: "WebKit", ClientEngineVersion: "605.1.15", DeviceName: "smartphone", DeviceBrand: "Google", DeviceModel: "Nexus 5", CountryCode: "US", CountryName: "United States", Current: true, Factors: []string{"test"}, Secret: "5e5bb8c16897e", MfaUpdatedAt: "2020-10-15T06:38:00.000+00:00"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result SessionList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestSessionListModelDefaults(t *testing.T) {
	var fresh SessionList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := SessionList{Total: 5, Sessions: []Session{Session{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5bb8c16897e", Expire: "2020-10-15T06:38:00.000+00:00", Provider: "email", ProviderUid: "user@example.com", ProviderAccessToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", ProviderAccessTokenExpiry: "2020-10-15T06:38:00.000+00:00", ProviderRefreshToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", Ip: "127.0.0.1", OsCode: "Mac", OsName: "Mac", OsVersion: "Mac", ClientType: "browser", ClientCode: "CM", ClientName: "Chrome Mobile iOS", ClientVersion: "84.0", ClientEngine: "WebKit", ClientEngineVersion: "605.1.15", DeviceName: "smartphone", DeviceBrand: "Google", DeviceModel: "Nexus 5", CountryCode: "US", CountryName: "United States", Current: true, Factors: []string{"test"}, Secret: "5e5bb8c16897e", MfaUpdatedAt: "2020-10-15T06:38:00.000+00:00"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
