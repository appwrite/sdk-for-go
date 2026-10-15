package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFrameworkListModel(t *testing.T) {
	model := FrameworkList{Total: 5, Frameworks: []Framework{Framework{Key: "sveltekit", Name: "SvelteKit", BuildRuntime: "node-22", Runtimes: []string{"test"}, Adapters: []FrameworkAdapter{FrameworkAdapter{Key: "static", InstallCommand: "npm install", BuildCommand: "npm run build", OutputDirectory: "./dist"}}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result FrameworkList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestFrameworkListModelDefaults(t *testing.T) {
	var fresh FrameworkList
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

	reused := FrameworkList{Total: 5, Frameworks: []Framework{Framework{Key: "sveltekit", Name: "SvelteKit", BuildRuntime: "node-22", Runtimes: []string{"test"}, Adapters: []FrameworkAdapter{FrameworkAdapter{Key: "static", InstallCommand: "npm install", BuildCommand: "npm run build", OutputDirectory: "./dist"}}}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
