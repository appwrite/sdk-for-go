package models

import (
    "encoding/json"
    "testing"
)

func TestLogModel(t *testing.T) {
    model := Log{        Event: "account.sessions.create",        UserId: "610fc2f985ee0",        UserEmail: "john@appwrite.io",        UserName: "John Doe",        Mode: "admin",        UserType: "user",        Ip: "127.0.0.1",        Time: "2020-10-15T06:38:00.000+00:00",        OsCode: "Mac",        OsName: "Mac",        OsVersion: "Mac",        ClientType: "browser",        ClientCode: "CM",        ClientName: "Chrome Mobile iOS",        ClientVersion: "84.0",        ClientEngine: "WebKit",        ClientEngineVersion: "605.1.15",        DeviceName: "smartphone",        DeviceBrand: "Google",        DeviceModel: "Nexus 5",        CountryCode: "US",        CountryName: "United States",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Log
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Event != model.Event {
        t.Errorf("Expected Event %v, got %v", model.Event, result.Event)
    }
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.UserEmail != model.UserEmail {
        t.Errorf("Expected UserEmail %v, got %v", model.UserEmail, result.UserEmail)
    }
    if result.UserName != model.UserName {
        t.Errorf("Expected UserName %v, got %v", model.UserName, result.UserName)
    }
    if result.Mode != model.Mode {
        t.Errorf("Expected Mode %v, got %v", model.Mode, result.Mode)
    }
    if result.UserType != model.UserType {
        t.Errorf("Expected UserType %v, got %v", model.UserType, result.UserType)
    }
    if result.Ip != model.Ip {
        t.Errorf("Expected Ip %v, got %v", model.Ip, result.Ip)
    }
    if result.Time != model.Time {
        t.Errorf("Expected Time %v, got %v", model.Time, result.Time)
    }
    if result.OsCode != model.OsCode {
        t.Errorf("Expected OsCode %v, got %v", model.OsCode, result.OsCode)
    }
    if result.OsName != model.OsName {
        t.Errorf("Expected OsName %v, got %v", model.OsName, result.OsName)
    }
    if result.OsVersion != model.OsVersion {
        t.Errorf("Expected OsVersion %v, got %v", model.OsVersion, result.OsVersion)
    }
    if result.ClientType != model.ClientType {
        t.Errorf("Expected ClientType %v, got %v", model.ClientType, result.ClientType)
    }
    if result.ClientCode != model.ClientCode {
        t.Errorf("Expected ClientCode %v, got %v", model.ClientCode, result.ClientCode)
    }
    if result.ClientName != model.ClientName {
        t.Errorf("Expected ClientName %v, got %v", model.ClientName, result.ClientName)
    }
    if result.ClientVersion != model.ClientVersion {
        t.Errorf("Expected ClientVersion %v, got %v", model.ClientVersion, result.ClientVersion)
    }
    if result.ClientEngine != model.ClientEngine {
        t.Errorf("Expected ClientEngine %v, got %v", model.ClientEngine, result.ClientEngine)
    }
    if result.ClientEngineVersion != model.ClientEngineVersion {
        t.Errorf("Expected ClientEngineVersion %v, got %v", model.ClientEngineVersion, result.ClientEngineVersion)
    }
    if result.DeviceName != model.DeviceName {
        t.Errorf("Expected DeviceName %v, got %v", model.DeviceName, result.DeviceName)
    }
    if result.DeviceBrand != model.DeviceBrand {
        t.Errorf("Expected DeviceBrand %v, got %v", model.DeviceBrand, result.DeviceBrand)
    }
    if result.DeviceModel != model.DeviceModel {
        t.Errorf("Expected DeviceModel %v, got %v", model.DeviceModel, result.DeviceModel)
    }
    if result.CountryCode != model.CountryCode {
        t.Errorf("Expected CountryCode %v, got %v", model.CountryCode, result.CountryCode)
    }
    if result.CountryName != model.CountryName {
        t.Errorf("Expected CountryName %v, got %v", model.CountryName, result.CountryName)
    }}
