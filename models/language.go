package models


// Language Model
type Language struct {
    // Language name.
    Name string `json:"name"`
    // Language two-character ISO 639-1 codes.
    Code string `json:"code"`
    // Language native name.
    NativeName string `json:"nativeName"`

}
