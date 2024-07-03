package models


// Phone Model
type Phone struct {
    // Phone code.
    Code string `json:"code"`
    // Country two-character ISO 3166-1 alpha code.
    CountryCode string `json:"countryCode"`
    // Country name.
    CountryName string `json:"countryName"`

}
