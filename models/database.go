package models


// Database Model
type Database struct {
    // Database ID.
    Id string `json:"$id"`
    // Database name.
    Name string `json:"name"`
    // Database creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Database update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // If database is enabled. Can be 'enabled' or 'disabled'. When disabled, the
    // database is inaccessible to users, but remains accessible to Server SDKs
    // using API keys.
    Enabled bool `json:"enabled"`

}
