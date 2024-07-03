package models


// FilesList Model
type FileList struct {
    // Total number of files documents that matched your query.
    Total int `json:"total"`
    // List of files.
    Files []interface{} `json:"files"`

}
