package models


// UsersList Model
type UserList struct {
    // Total number of users documents that matched your query.
    Total int `json:"total"`
    // List of users.
    Users []interface{} `json:"users"`

}
