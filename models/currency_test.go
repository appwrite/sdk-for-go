package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCurrencyModel(t *testing.T) {
	model := Currency{Symbol: "$", Name: "US dollar", SymbolNative: "$", DecimalDigits: 2, Rounding: 0, Code: "USD", NamePlural: "US dollars"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Currency
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Symbol != model.Symbol {
		t.Errorf("Expected Symbol %v, got %v", model.Symbol, result.Symbol)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.SymbolNative != model.SymbolNative {
		t.Errorf("Expected SymbolNative %v, got %v", model.SymbolNative, result.SymbolNative)
	}
	if result.DecimalDigits != model.DecimalDigits {
		t.Errorf("Expected DecimalDigits %v, got %v", model.DecimalDigits, result.DecimalDigits)
	}
	if result.Rounding != model.Rounding {
		t.Errorf("Expected Rounding %v, got %v", model.Rounding, result.Rounding)
	}
	if result.Code != model.Code {
		t.Errorf("Expected Code %v, got %v", model.Code, result.Code)
	}
	if result.NamePlural != model.NamePlural {
		t.Errorf("Expected NamePlural %v, got %v", model.NamePlural, result.NamePlural)
	}
}

func TestCurrencyModelDefaults(t *testing.T) {
	var fresh Currency
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Symbol)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing symbol to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Name)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing name to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.SymbolNative)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing symbolNative to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.DecimalDigits)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing decimalDigits to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Rounding)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing rounding to default to %v, got %v", want, got)
		}
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
		encoded, err := json.Marshal(fresh.NamePlural)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing namePlural to default to %v, got %v", want, got)
		}
	}

	reused := Currency{Symbol: "$", Name: "US dollar", SymbolNative: "$", DecimalDigits: 2, Rounding: 0, Code: "USD", NamePlural: "US dollars"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
