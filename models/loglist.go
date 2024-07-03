package models


// LogsList Model
type LogList struct {
    // Total number of logs documents that matched your query.
    Total int `json:"total"`
    // List of logs.
    Logs []interface{} `json:"logs"`

}
