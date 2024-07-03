package models


// PhonesList Model
type PhoneList struct {
    // Total number of phones documents that matched your query.
    Total int `json:"total"`
    // List of phones.
    Phones []interface{} `json:"phones"`

}
