package models


// CountriesList Model
type CountryList struct {
    // Total number of countries documents that matched your query.
    Total int `json:"total"`
    // List of countries.
    Countries []interface{} `json:"countries"`

}
