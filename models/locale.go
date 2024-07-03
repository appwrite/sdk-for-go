package models


// Locale Model
type Locale struct {
    // User IP address.
    Ip string `json:"ip"`
    // Country code in [ISO 3166-1](http://en.wikipedia.org/wiki/ISO_3166-1)
    // two-character format
    CountryCode string `json:"countryCode"`
    // Country name. This field support localization.
    Country string `json:"country"`
    // Continent code. A two character continent code "AF" for Africa, "AN" for
    // Antarctica, "AS" for Asia, "EU" for Europe, "NA" for North America, "OC"
    // for Oceania, and "SA" for South America.
    ContinentCode string `json:"continentCode"`
    // Continent name. This field support localization.
    Continent string `json:"continent"`
    // True if country is part of the European Union.
    Eu bool `json:"eu"`
    // Currency code in [ISO 4217-1](http://en.wikipedia.org/wiki/ISO_4217)
    // three-character format
    Currency string `json:"currency"`

}
