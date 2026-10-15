package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCurrencyListModel(t *testing.T) {
	model := CurrencyList{Total: 5, Currencies: []Currency{Currency{Symbol: "$", Name: "US dollar", SymbolNative: "$", DecimalDigits: 2, Rounding: 0, Code: "USD", NamePlural: "US dollars"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result CurrencyList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestCurrencyListModelDefaults(t *testing.T) {
	var fresh CurrencyList
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Total)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing total to default to %v, got %v", want, got)
		}
	}

	reused := CurrencyList{Total: 5, Currencies: []Currency{Currency{Symbol: "$", Name: "US dollar", SymbolNative: "$", DecimalDigits: 2, Rounding: 0, Code: "USD", NamePlural: "US dollars"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
