package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAlgoPhpassModel(t *testing.T) {
	model := AlgoPhpass{Type: "phpass"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AlgoPhpass
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
}

func TestAlgoPhpassModelDefaults(t *testing.T) {
	var fresh AlgoPhpass
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"phpass\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Type)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing type to default to %v, got %v", want, got)
		}
	}

	reused := AlgoPhpass{Type: "phpass"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
