package models

import (
    "encoding/json"
    "testing"
)

func TestUserListModel(t *testing.T) {
    model := UserList{        Total: 5,        Users: []User{User{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "John Doe",        Registration: "2020-10-15T06:38:00.000+00:00",        Status: true,        Labels: []string{"test"},        PasswordUpdate: "2020-10-15T06:38:00.000+00:00",        Email: "john@appwrite.io",        Phone: "+4930901820",        EmailVerification: true,        PhoneVerification: true,        Mfa: true,        Prefs: Preferences{    },        Targets: []Target{Target{        Id: "259125845563242502",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        Name: "Apple iPhone 12",        UserId: "259125845563242502",        ProviderType: "email",        Identifier: "token",        Expired: true,    },
            },        AccessedAt: "2020-10-15T06:38:00.000+00:00",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result UserList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
