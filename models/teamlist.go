package models


// TeamsList Model
type TeamList struct {
    // Total number of teams documents that matched your query.
    Total int `json:"total"`
    // List of teams.
    Teams []interface{} `json:"teams"`

}
