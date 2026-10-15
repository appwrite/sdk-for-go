package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAlgoMd5Model(t *testing.T) {
	model := AlgoMd5{Type: "md5"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AlgoMd5
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
}

func TestAlgoMd5ModelDefaults(t *testing.T) {
	var fresh AlgoMd5
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"md5\""), &want); err != nil {
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

	reused := AlgoMd5{Type: "md5"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
