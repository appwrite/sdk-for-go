package models


// ExecutionsList Model
type ExecutionList struct {
    // Total number of executions documents that matched your query.
    Total int `json:"total"`
    // List of executions.
    Executions []interface{} `json:"executions"`

}
