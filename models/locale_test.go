package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLocaleModel(t *testing.T) {
	model := Locale{Ip: "127.0.0.1", CountryCode: "US", Country: "United States", ContinentCode: "NA", Continent: "North America", Eu: true, Currency: "USD"}

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
	}
}

func TestLocaleModelDefaults(t *testing.T) {
	var fresh Locale
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Ip)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing ip to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CountryCode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing countryCode to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Country)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing country to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.ContinentCode)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing continentCode to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Continent)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing continent to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Eu)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing eu to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Currency)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing currency to default to %v, got %v", want, got)
		}
	}

	reused := Locale{Ip: "127.0.0.1", CountryCode: "US", Country: "United States", ContinentCode: "NA", Continent: "North America", Eu: true, Currency: "USD"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
