package models


// Log Model
type Log struct {
    // Event name.
    Event string `json:"event"`
    // User ID.
    UserId string `json:"userId"`
    // User Email.
    UserEmail string `json:"userEmail"`
    // User Name.
    UserName string `json:"userName"`
    // API mode when event triggered.
    Mode string `json:"mode"`
    // IP session in use when the session was created.
    Ip string `json:"ip"`
    // Log creation date in ISO 8601 format.
    Time string `json:"time"`
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

}
