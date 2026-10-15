package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestIdentityListModel(t *testing.T) {
	model := IdentityList{Total: 5, Identities: []Identity{Identity{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5bb8c16897e", Provider: "email", ProviderUid: "5e5bb8c16897e", ProviderEmail: "user@example.com", ProviderAccessToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", ProviderAccessTokenExpiry: "2020-10-15T06:38:00.000+00:00", ProviderRefreshToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", ProviderIdToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBhYzNmMWQwNWExYjhlN2YifQ.eyJzdWIiOiIxMTAxNjk0ODQ0NzQzODYyNzYzMzQifQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result IdentityList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestIdentityListModelDefaults(t *testing.T) {
	var fresh IdentityList
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

	reused := IdentityList{Total: 5, Identities: []Identity{Identity{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5bb8c16897e", Provider: "email", ProviderUid: "5e5bb8c16897e", ProviderEmail: "user@example.com", ProviderAccessToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", ProviderAccessTokenExpiry: "2020-10-15T06:38:00.000+00:00", ProviderRefreshToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3", ProviderIdToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBhYzNmMWQwNWExYjhlN2YifQ.eyJzdWIiOiIxMTAxNjk0ODQ0NzQzODYyNzYzMzQifQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
