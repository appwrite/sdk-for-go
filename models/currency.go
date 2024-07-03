package models


// Currency Model
type Currency struct {
    // Currency symbol.
    Symbol string `json:"symbol"`
    // Currency name.
    Name string `json:"name"`
    // Currency native symbol.
    SymbolNative string `json:"symbolNative"`
    // Number of decimal digits.
    DecimalDigits int `json:"decimalDigits"`
    // Currency digit rounding.
    Rounding float64 `json:"rounding"`
    // Currency code in [ISO 4217-1](http://en.wikipedia.org/wiki/ISO_4217)
    // three-character format.
    Code string `json:"code"`
    // Currency plural name
    NamePlural string `json:"namePlural"`

}
