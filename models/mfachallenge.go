package models


// MFAChallenge Model
type MfaChallenge struct {
    // Token ID.
    Id string `json:"$id"`
    // Token creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // User ID.
    UserId string `json:"userId"`
    // Token expiration date in ISO 8601 format.
    Expire string `json:"expire"`

}
