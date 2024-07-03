package models


// ContinentsList Model
type ContinentList struct {
    // Total number of continents documents that matched your query.
    Total int `json:"total"`
    // List of continents.
    Continents []interface{} `json:"continents"`

}
