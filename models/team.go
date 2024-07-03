package models


// Team Model
type Team struct {
    // Team ID.
    Id string `json:"$id"`
    // Team creation date in ISO 8601 format.
    CreatedAt string `json:"$createdAt"`
    // Team update date in ISO 8601 format.
    UpdatedAt string `json:"$updatedAt"`
    // Team name.
    Name string `json:"name"`
    // Total number of team members.
    Total int `json:"total"`
    // Team preferences as a key-value object
    Prefs interface{} `json:"prefs"`

}
