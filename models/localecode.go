package models


// LocaleCode Model
type LocaleCode struct {
    // Locale codes in [ISO
    // 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes)
    Code string `json:"code"`
    // Locale name
    Name string `json:"name"`

}
