package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSpecificationModel(t *testing.T) {
	model := Specification{Memory: 512, Cpus: 1, Enabled: true, Slug: "s-1vcpu-512mb"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Specification
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Memory != model.Memory {
		t.Errorf("Expected Memory %v, got %v", model.Memory, result.Memory)
	}
	if result.Cpus != model.Cpus {
		t.Errorf("Expected Cpus %v, got %v", model.Cpus, result.Cpus)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
	if result.Slug != model.Slug {
		t.Errorf("Expected Slug %v, got %v", model.Slug, result.Slug)
	}
}

func TestSpecificationModelDefaults(t *testing.T) {
	var fresh Specification
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Memory)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing memory to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Cpus)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing cpus to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("false"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Enabled)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing enabled to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Slug)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing slug to default to %v, got %v", want, got)
		}
	}

	reused := Specification{Memory: 512, Cpus: 1, Enabled: true, Slug: "s-1vcpu-512mb"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
