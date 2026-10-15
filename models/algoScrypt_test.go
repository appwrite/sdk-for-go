package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAlgoScryptModel(t *testing.T) {
	model := AlgoScrypt{Type: "scrypt", CostCpu: 8, CostMemory: 14, CostParallel: 1, Length: 64}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AlgoScrypt
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Type != model.Type {
		t.Errorf("Expected Type %v, got %v", model.Type, result.Type)
	}
	if result.CostCpu != model.CostCpu {
		t.Errorf("Expected CostCpu %v, got %v", model.CostCpu, result.CostCpu)
	}
	if result.CostMemory != model.CostMemory {
		t.Errorf("Expected CostMemory %v, got %v", model.CostMemory, result.CostMemory)
	}
	if result.CostParallel != model.CostParallel {
		t.Errorf("Expected CostParallel %v, got %v", model.CostParallel, result.CostParallel)
	}
	if result.Length != model.Length {
		t.Errorf("Expected Length %v, got %v", model.Length, result.Length)
	}
}

func TestAlgoScryptModelDefaults(t *testing.T) {
	var fresh AlgoScrypt
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"scrypt\""), &want); err != nil {
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
		if err := json.Unmarshal([]byte("8"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CostCpu)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing costCpu to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("14"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CostMemory)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing costMemory to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("1"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.CostParallel)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing costParallel to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("64"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Length)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing length to default to %v, got %v", want, got)
		}
	}

	reused := AlgoScrypt{Type: "scrypt", CostCpu: 8, CostMemory: 14, CostParallel: 1, Length: 64}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
