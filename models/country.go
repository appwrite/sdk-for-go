package models


// Country Model
type Country struct {
    // Country name.
    Name string `json:"name"`
    // Country two-character ISO 3166-1 alpha code.
    Code string `json:"code"`

}
