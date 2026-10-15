package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFrameworkModel(t *testing.T) {
	model := Framework{Key: "sveltekit", Name: "SvelteKit", BuildRuntime: "node-22", Runtimes: []string{"test"}, Adapters: []FrameworkAdapter{FrameworkAdapter{Key: "static", InstallCommand: "npm install", BuildCommand: "npm run build", OutputDirectory: "./dist"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result Framework
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Key != model.Key {
		t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.BuildRuntime != model.BuildRuntime {
		t.Errorf("Expected BuildRuntime %v, got %v", model.BuildRuntime, result.BuildRuntime)
	}
}

func TestFrameworkModelDefaults(t *testing.T) {
	var fresh Framework
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Key)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing key to default to %v, got %v", want, got)
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
		encoded, err := json.Marshal(fresh.BuildRuntime)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing buildRuntime to default to %v, got %v", want, got)
		}
	}

	reused := Framework{Key: "sveltekit", Name: "SvelteKit", BuildRuntime: "node-22", Runtimes: []string{"test"}, Adapters: []FrameworkAdapter{FrameworkAdapter{Key: "static", InstallCommand: "npm install", BuildCommand: "npm run build", OutputDirectory: "./dist"}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
