package models


// LocaleCodesList Model
type LocaleCodeList struct {
    // Total number of localeCodes documents that matched your query.
    Total int `json:"total"`
    // List of localeCodes.
    LocaleCodes []interface{} `json:"localeCodes"`

}
