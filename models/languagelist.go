package models


// LanguagesList Model
type LanguageList struct {
    // Total number of languages documents that matched your query.
    Total int `json:"total"`
    // List of languages.
    Languages []interface{} `json:"languages"`

}
