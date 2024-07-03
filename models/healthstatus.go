package models


// HealthStatus Model
type HealthStatus struct {
    // Name of the service.
    Name string `json:"name"`
    // Duration in milliseconds how long the health check took.
    Ping int `json:"ping"`
    // Service status. Possible values can are: `pass`, `fail`
    Status string `json:"status"`

}
