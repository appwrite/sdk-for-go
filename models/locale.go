package models

import (
    "encoding/json"
    "errors"
)

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
    // City
    City string `json:"city"`
    // Name of timezone
    TimeZone string `json:"timeZone"`
    // Postal code
    PostalCode string `json:"postalCode"`
    // Latitude
    Latitude float64 `json:"latitude"`
    // Longitude
    Longitude float64 `json:"longitude"`
    // Autonomous System Number (ASN) of the IP
    AutonomousSystemNumber string `json:"autonomousSystemNumber"`
    // Organization that owns the ASN
    AutonomousSystemOrganization string `json:"autonomousSystemOrganization"`
    // Internet service provider of the IP
    Isp string `json:"isp"`
    // Connection type of the IP (e.g. cable, cellular, corporate)
    ConnectionType string `json:"connectionType"`
    // User type classification of the IP (e.g. residential, business, hosting)
    ConnectionUsageType string `json:"connectionUsageType"`
    // Registered organization of the IP
    ConnectionOrganization string `json:"connectionOrganization"`

    // Used by Decode() method
    data []byte
}

func (model Locale) New(data []byte) *Locale {
    model.data = data
    return &model
}

func (model *Locale) Decode(value interface{}) error {
    if len(model.data) <= 0 {
        return errors.New("method Decode() cannot be used on nested struct")
    }

    err := json.Unmarshal(model.data, value)
    if err != nil {
        return err
    }

    return nil
}