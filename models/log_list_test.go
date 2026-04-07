package models

import (
    "encoding/json"
    "testing"
)

func TestLogListModel(t *testing.T) {
    model := LogList{        Total: 5,        Logs: []Log{Log{        Event: "account.sessions.create",        UserId: "610fc2f985ee0",        UserEmail: "john@appwrite.io",        UserName: "John Doe",        Mode: "admin",        Ip: "127.0.0.1",        Time: "2020-10-15T06:38:00.000+00:00",        OsCode: "Mac",        OsName: "Mac",        OsVersion: "Mac",        ClientType: "browser",        ClientCode: "CM",        ClientName: "Chrome Mobile iOS",        ClientVersion: "84.0",        ClientEngine: "WebKit",        ClientEngineVersion: "605.1.15",        DeviceName: "smartphone",        DeviceBrand: "Google",        DeviceModel: "Nexus 5",        CountryCode: "US",        CountryName: "United States",    },
            },    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result LogList
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Total != model.Total {
        t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
    }}
