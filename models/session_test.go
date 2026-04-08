package models

import (
    "encoding/json"
    "testing"
)

func TestSessionModel(t *testing.T) {
    model := Session{        Id: "5e5ea5c16897e",        CreatedAt: "2020-10-15T06:38:00.000+00:00",        UpdatedAt: "2020-10-15T06:38:00.000+00:00",        UserId: "5e5bb8c16897e",        Expire: "2020-10-15T06:38:00.000+00:00",        Provider: "email",        ProviderUid: "user@example.com",        ProviderAccessToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",        ProviderAccessTokenExpiry: "2020-10-15T06:38:00.000+00:00",        ProviderRefreshToken: "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3",        Ip: "127.0.0.1",        OsCode: "Mac",        OsName: "Mac",        OsVersion: "Mac",        ClientType: "browser",        ClientCode: "CM",        ClientName: "Chrome Mobile iOS",        ClientVersion: "84.0",        ClientEngine: "WebKit",        ClientEngineVersion: "605.1.15",        DeviceName: "smartphone",        DeviceBrand: "Google",        DeviceModel: "Nexus 5",        CountryCode: "US",        CountryName: "United States",        Current: true,        Factors: []string{"test"},        Secret: "5e5bb8c16897e",        MfaUpdatedAt: "2020-10-15T06:38:00.000+00:00",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Session
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
    if result.UserId != model.UserId {
        t.Errorf("Expected UserId %v, got %v", model.UserId, result.UserId)
    }
    if result.Expire != model.Expire {
        t.Errorf("Expected Expire %v, got %v", model.Expire, result.Expire)
    }
    if result.Provider != model.Provider {
        t.Errorf("Expected Provider %v, got %v", model.Provider, result.Provider)
    }
    if result.ProviderUid != model.ProviderUid {
        t.Errorf("Expected ProviderUid %v, got %v", model.ProviderUid, result.ProviderUid)
    }
    if result.ProviderAccessToken != model.ProviderAccessToken {
        t.Errorf("Expected ProviderAccessToken %v, got %v", model.ProviderAccessToken, result.ProviderAccessToken)
    }
    if result.ProviderAccessTokenExpiry != model.ProviderAccessTokenExpiry {
        t.Errorf("Expected ProviderAccessTokenExpiry %v, got %v", model.ProviderAccessTokenExpiry, result.ProviderAccessTokenExpiry)
    }
    if result.ProviderRefreshToken != model.ProviderRefreshToken {
        t.Errorf("Expected ProviderRefreshToken %v, got %v", model.ProviderRefreshToken, result.ProviderRefreshToken)
    }
    if result.Ip != model.Ip {
        t.Errorf("Expected Ip %v, got %v", model.Ip, result.Ip)
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
    }
    if result.Current != model.Current {
        t.Errorf("Expected Current %v, got %v", model.Current, result.Current)
    }
    if result.Secret != model.Secret {
        t.Errorf("Expected Secret %v, got %v", model.Secret, result.Secret)
    }
    if result.MfaUpdatedAt != model.MfaUpdatedAt {
        t.Errorf("Expected MfaUpdatedAt %v, got %v", model.MfaUpdatedAt, result.MfaUpdatedAt)
    }}
