package models

import (
    "encoding/json"
    "testing"
)

func TestActivityEventListModel(t *testing.T) {
    model := ActivityEventList{        Total: 5,        Events: []ActivityEvent{ActivityEvent{        Id: "5e5ea5c16897e",        UserType: "user",        UserId: "610fc2f985ee0",        UserEmail: "john@appwrite.io",        UserName: "John Doe",        ResourceParent: "database/ID",        ResourceType: "collection",        ResourceId: "610fc2f985ee0",        Resource: "collections/610fc2f985ee0",        Event: "account.sessions.create",        UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",        Ip: "127.0.0.1",        Mode: "admin",        Country: "US",        Time: "2020-10-15T06:38:00.000+00:00",        ProjectId: "610fc2f985ee0",        TeamId: "610fc2f985ee0",        Hostname: "appwrite.io",        OsCode: "Mac",        OsName: "Mac",        OsVersion: "Mac",        ClientType: "browser",        ClientCode: "CM",        ClientName: "Chrome Mobile iOS",        ClientVersion: "84.0",        ClientEngine: "WebKit",        ClientEngineVersion: "605.1.15",        DeviceName: "smartphone",        DeviceBrand: "Google",        DeviceModel: "Nexus 5",        CountryCode: "US",        CountryName: "United States",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result ActivityEventList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
