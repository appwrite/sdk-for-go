package models


// CurrenciesList Model
type CurrencyList struct {
    // Total number of currencies documents that matched your query.
    Total int `json:"total"`
    // List of currencies.
    Currencies []interface{} `json:"currencies"`

}
