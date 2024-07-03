package models


// Target Model
type Target struct {
    // Target ID.
    Id string `json:"$id"`
    // Target creation time in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Target update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Target Name.
    Name string `json:"name"`
    // User ID.
    UserId string `json:"userId"`
    // Provider ID.
    ProviderId string `json:"providerId"`
    // The target provider type. Can be one of the following: `email`, `sms` or
    // `push`.
    ProviderType string `json:"providerType"`
    // The target identifier.
    Identifier string `json:"identifier"`

}
