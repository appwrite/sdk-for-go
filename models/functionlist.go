package models


// FunctionsList Model
type FunctionList struct {
    // Total number of functions documents that matched your query.
    Total int `json:"total"`
    // List of functions.
    Functions []interface{} `json:"functions"`

}
