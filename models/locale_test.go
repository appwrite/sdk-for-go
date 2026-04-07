package models

import (
    "encoding/json"
    "testing"
)

func TestLocaleModel(t *testing.T) {
    model := Locale{        Ip: "127.0.0.1",        CountryCode: "US",        Country: "United States",        ContinentCode: "NA",        Continent: "North America",        Eu: true,        Currency: "USD",    }

    data, err := json.Marshal(model)
    if err != nil {
        t.Fatal(err)
    }

    var result Locale
    err = json.Unmarshal(data, &result)
    if err != nil {
        t.Fatal(err)
    }
    if result.Ip != model.Ip {
        t.Errorf("Expected Ip %v, got %v", model.Ip, result.Ip)
    }
    if result.CountryCode != model.CountryCode {
        t.Errorf("Expected CountryCode %v, got %v", model.CountryCode, result.CountryCode)
    }
    if result.Country != model.Country {
        t.Errorf("Expected Country %v, got %v", model.Country, result.Country)
    }
    if result.ContinentCode != model.ContinentCode {
        t.Errorf("Expected ContinentCode %v, got %v", model.ContinentCode, result.ContinentCode)
    }
    if result.Continent != model.Continent {
        t.Errorf("Expected Continent %v, got %v", model.Continent, result.Continent)
    }
    if result.Eu != model.Eu {
        t.Errorf("Expected Eu %v, got %v", model.Eu, result.Eu)
    }
    if result.Currency != model.Currency {
        t.Errorf("Expected Currency %v, got %v", model.Currency, result.Currency)
    }}
