package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAlgoArgon2Model(t *testing.T) {
	model := AlgoArgon2{Type: "argon2", MemoryCost: 65536, TimeCost: 4, Threads: 3}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AlgoArgon2
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
	if result.MemoryCost != model.MemoryCost {
		t.Errorf("Expected MemoryCost %v, got %v", model.MemoryCost, result.MemoryCost)
	}
	if result.TimeCost != model.TimeCost {
		t.Errorf("Expected TimeCost %v, got %v", model.TimeCost, result.TimeCost)
	}
	if result.Threads != model.Threads {
		t.Errorf("Expected Threads %v, got %v", model.Threads, result.Threads)
	}
}

func TestAlgoArgon2ModelDefaults(t *testing.T) {
	var fresh AlgoArgon2
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"argon2\""), &want); err != nil {
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
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.MemoryCost)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing memoryCost to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.TimeCost)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing timeCost to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Threads)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing threads to default to %v, got %v", want, got)
		}
	}

	reused := AlgoArgon2{Type: "argon2", MemoryCost: 65536, TimeCost: 4, Threads: 3}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
