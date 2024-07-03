package models


// MembershipsList Model
type MembershipList struct {
    // Total number of memberships documents that matched your query.
    Total int `json:"total"`
    // List of memberships.
    Memberships []interface{} `json:"memberships"`

}
