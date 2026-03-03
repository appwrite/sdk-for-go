package models

import (
    "encoding/json"
    "errors"
)

// ActivityEvent Model
type ActivityEvent struct {
    // Event ID.
    Id string `json:"$id"`
    // User type.
    UserType string `json:"userType"`
    // User ID.
    UserId string `json:"userId"`
    // User Email.
    UserEmail string `json:"userEmail"`
    // User Name.
    UserName string `json:"userName"`
    // Resource parent.
    ResourceParent string `json:"resourceParent"`
    // Resource type.
    ResourceType string `json:"resourceType"`
    // Resource ID.
    ResourceId string `json:"resourceId"`
    // Resource.
    Resource string `json:"resource"`
    // Event name.
    Event string `json:"event"`
    // User agent.
    UserAgent string `json:"userAgent"`
    // IP address.
    Ip string `json:"ip"`
    // API mode when event triggered.
    Mode string `json:"mode"`
    // Location.
    Country string `json:"country"`
    // Log creation date in ISO 8601 format.
    Time string `json:"time"`
    // Project ID.
    ProjectId string `json:"projectId"`
    // Team ID.
    TeamId string `json:"teamId"`
    // Hostname.
    Hostname string `json:"hostname"`
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

    // Used by Decode() method
    data []byte
}

func (model ActivityEvent) New(data []byte) *ActivityEvent {
    model.data = data
    return &model
}

func (model *ActivityEvent) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}