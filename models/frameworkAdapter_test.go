package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFrameworkAdapterModel(t *testing.T) {
	model := FrameworkAdapter{Key: "static", InstallCommand: "npm install", BuildCommand: "npm run build", OutputDirectory: "./dist"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result FrameworkAdapter
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Key != model.Key {
		t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
	}
	if result.InstallCommand != model.InstallCommand {
		t.Errorf("Expected InstallCommand %v, got %v", model.InstallCommand, result.InstallCommand)
	}
	if result.BuildCommand != model.BuildCommand {
		t.Errorf("Expected BuildCommand %v, got %v", model.BuildCommand, result.BuildCommand)
	}
	if result.OutputDirectory != model.OutputDirectory {
		t.Errorf("Expected OutputDirectory %v, got %v", model.OutputDirectory, result.OutputDirectory)
	}
}

func TestFrameworkAdapterModelDefaults(t *testing.T) {
	var fresh FrameworkAdapter
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
		encoded, err := json.Marshal(fresh.InstallCommand)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing installCommand to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.BuildCommand)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing buildCommand to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.OutputDirectory)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing outputDirectory to default to %v, got %v", want, got)
		}
	}

	reused := FrameworkAdapter{Key: "static", InstallCommand: "npm install", BuildCommand: "npm run build", OutputDirectory: "./dist"}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
