package models


// RuntimesList Model
type RuntimeList struct {
    // Total number of runtimes documents that matched your query.
    Total int `json:"total"`
    // List of runtimes.
    Runtimes []interface{} `json:"runtimes"`

}
