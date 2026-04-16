package models

import (
    "encoding/json"
    "testing"
)

func TestMembershipListModel(t *testing.T) {
    model := MembershipList{        Total: 5,        Memberships: []Membership{Membership{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        UserId: "5e5ea5c16897e",        UserName: "John Doe",        UserEmail: "john@appwrite.io",        TeamId: "5e5ea5c16897e",        TeamName: "VIP",        Invited: "2020-10-15T06:38:00.000+00:00",        Joined: "2020-10-15T06:38:00.000+00:00",        Confirm: true,        Mfa: true,        Roles: []string{"test"},    },
            },    }

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
    }}
