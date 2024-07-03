package models


// Membership Model
type Membership struct {
    // Membership ID.
    Id string `json:"$id"`
    // Membership creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Membership update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // User ID.
    UserId string `json:"userId"`
    // User name.
    UserName string `json:"userName"`
    // User email address.
    UserEmail string `json:"userEmail"`
    // Team ID.
    TeamId string `json:"teamId"`
    // Team name.
    TeamName string `json:"teamName"`
    // Date, the user has been invited to join the team in ISO 8601 format.
    Invited string `json:"invited"`
    // Date, the user has accepted the invitation to join the team in ISO 8601
    // format.
    Joined string `json:"joined"`
    // User confirmation status, true if the user has joined the team or false
    // otherwise.
    Confirm bool `json:"confirm"`
    // Multi factor authentication status, true if the user has MFA enabled or
    // false otherwise.
    Mfa bool `json:"mfa"`
    // User list of roles
    Roles []interface{} `json:"roles"`

}
