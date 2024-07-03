package models


// IdentitiesList Model
type IdentityList struct {
    // Total number of identities documents that matched your query.
    Total int `json:"total"`
    // List of identities.
    Identities []interface{} `json:"identities"`

}
