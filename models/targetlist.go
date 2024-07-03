package models


// TargetList Model
type TargetList struct {
    // Total number of targets documents that matched your query.
    Total int `json:"total"`
    // List of targets.
    Targets []interface{} `json:"targets"`

}
