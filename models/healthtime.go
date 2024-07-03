package models


// HealthTime Model
type HealthTime struct {
    // Current unix timestamp on trustful remote server.
    RemoteTime int `json:"remoteTime"`
    // Current unix timestamp of local server where Appwrite runs.
    LocalTime int `json:"localTime"`
    // Difference of unix remote and local timestamps in milliseconds.
    Diff int `json:"diff"`

}
