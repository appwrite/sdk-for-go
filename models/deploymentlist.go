package models


// DeploymentsList Model
type DeploymentList struct {
    // Total number of deployments documents that matched your query.
    Total int `json:"total"`
    // List of deployments.
    Deployments []interface{} `json:"deployments"`

}
