package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPhoneModel(t *testing.T) {
	model := Phone{Code: "+1", CountryCode: "US", CountryName: "United States"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Phone
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != model.Code {
		t.Errorf("Expected Code %v, got %v", model.Code, result.Code)
	}
	if result.CountryCode != model.CountryCode {
		t.Errorf("Expected CountryCode %v, got %v", model.CountryCode, result.CountryCode)
	}
	if result.CountryName != model.CountryName {
		t.Errorf("Expected CountryName %v, got %v", model.CountryName, result.CountryName)
	}
}

func TestPhoneModelDefaults(t *testing.T) {
	var fresh Phone
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Code)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing code to default to %v, got %v", want, got)
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
		encoded, err := json.Marshal(fresh.CountryName)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing countryName to default to %v, got %v", want, got)
		}
	}

	reused := Phone{Code: "+1", CountryCode: "US", CountryName: "United States"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
