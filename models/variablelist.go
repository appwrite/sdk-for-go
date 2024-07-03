package models


// VariablesList Model
type VariableList struct {
    // Total number of variables documents that matched your query.
    Total int `json:"total"`
    // List of variables.
    Variables []interface{} `json:"variables"`

}
