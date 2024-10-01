package models

import (
    "encoding/json"
    "errors"
)

// Session Model
type Session struct {
    // Session ID.
    Id string `json:"$id"`
    // Session creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Session update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // User ID.
    UserId string `json:"userId"`
    // Session expiration date in ISO 8601 format.
    Expire string `json:"expire"`
    // Session Provider.
    Provider string `json:"provider"`
    // Session Provider User ID.
    ProviderUid string `json:"providerUid"`
    // Session Provider Access Token.
    ProviderAccessToken string `json:"providerAccessToken"`
    // The date of when the access token expires in ISO 8601 format.
    ProviderAccessTokenExpiry string `json:"providerAccessTokenExpiry"`
    // Session Provider Refresh Token.
    ProviderRefreshToken string `json:"providerRefreshToken"`
    // IP in use when the session was created.
    Ip string `json:"ip"`
    // Operating system code name. View list of [available
    // options](https://github.com/appwrite/appwrite/blob/master/docs/lists/os.json).
    OsCode string `json:"osCode"`
    // Operating system name.
    OsName string `json:"osName"`
    // Operating system version.
    OsVersion string `json:"osVersion"`
    // Client type.
    ClientType string `json:"clientType"`
    // Client code name. View list of [available
    // options](https://github.com/appwrite/appwrite/blob/master/docs/lists/clients.json).
    ClientCode string `json:"clientCode"`
    // Client name.
    ClientName string `json:"clientName"`
    // Client version.
    ClientVersion string `json:"clientVersion"`
    // Client engine name.
    ClientEngine string `json:"clientEngine"`
    // Client engine name.
    ClientEngineVersion string `json:"clientEngineVersion"`
    // Device name.
    DeviceName string `json:"deviceName"`
    // Device brand name.
    DeviceBrand string `json:"deviceBrand"`
    // Device model name.
    DeviceModel string `json:"deviceModel"`
    // Country two-character ISO 3166-1 alpha code.
    CountryCode string `json:"countryCode"`
    // Country name.
    CountryName string `json:"countryName"`
    // Returns true if this the current user session.
    Current bool `json:"current"`
    // Returns a list of active session factors.
    Factors []string `json:"factors"`
    // Secret used to authenticate the user. Only included if the request was made
    // with an API key
    Secret string `json:"secret"`
    // Most recent date in ISO 8601 format when the session successfully passed
    // MFA challenge.
    MfaUpdatedAt string `json:"mfaUpdatedAt"`

    // Used by Decode() method
    data []byte
}

func (model Session) New(data []byte) *Session {
    model.data = data
    return &model
}

func (model *Session) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}
