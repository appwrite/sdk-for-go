package models


// Identity Model
type Identity struct {
    // Identity ID.
    Id string `json:"$id"`
    // Identity creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Identity update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // User ID.
    UserId string `json:"userId"`
    // Identity Provider.
    Provider string `json:"provider"`
    // ID of the User in the Identity Provider.
    ProviderUid string `json:"providerUid"`
    // Email of the User in the Identity Provider.
    ProviderEmail string `json:"providerEmail"`
    // Identity Provider Access Token.
    ProviderAccessToken string `json:"providerAccessToken"`
    // The date of when the access token expires in ISO 8601 format.
    ProviderAccessTokenExpiry string `json:"providerAccessTokenExpiry"`
    // Identity Provider Refresh Token.
    ProviderRefreshToken string `json:"providerRefreshToken"`

}
