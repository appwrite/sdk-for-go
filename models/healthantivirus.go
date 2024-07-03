package models


// HealthAntivirus Model
type HealthAntivirus struct {
    // Antivirus version.
    Version string `json:"version"`
    // Antivirus status. Possible values can are: `disabled`, `offline`, `online`
    Status string `json:"status"`

}
