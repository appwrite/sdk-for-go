package models


// IndexesList Model
type IndexList struct {
    // Total number of indexes documents that matched your query.
    Total int `json:"total"`
    // List of indexes.
    Indexes []interface{} `json:"indexes"`

}
