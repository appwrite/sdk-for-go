package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMembershipListModel(t *testing.T) {
	model := MembershipList{Total: 5, Memberships: []Membership{Membership{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5ea5c16897e", UserName: "John Doe", UserEmail: "john@appwrite.io", UserPhone: "+1 555 555 5555", TeamId: "5e5ea5c16897e", TeamName: "VIP", Invited: "2020-10-15T06:38:00.000+00:00", Joined: "2020-10-15T06:38:00.000+00:00", Confirm: true, Mfa: true, UserAccessedAt: "2020-10-15T06:38:00.000+00:00", Roles: []string{"test"}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result MembershipList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestMembershipListModelDefaults(t *testing.T) {
	var fresh MembershipList
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

	reused := MembershipList{Total: 5, Memberships: []Membership{Membership{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", UserId: "5e5ea5c16897e", UserName: "John Doe", UserEmail: "john@appwrite.io", UserPhone: "+1 555 555 5555", TeamId: "5e5ea5c16897e", TeamName: "VIP", Invited: "2020-10-15T06:38:00.000+00:00", Joined: "2020-10-15T06:38:00.000+00:00", Confirm: true, Mfa: true, UserAccessedAt: "2020-10-15T06:38:00.000+00:00", Roles: []string{"test"}}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
