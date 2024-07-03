package models


// ProviderList Model
type ProviderList struct {
    // Total number of providers documents that matched your query.
    Total int `json:"total"`
    // List of providers.
    Providers []interface{} `json:"providers"`

}
