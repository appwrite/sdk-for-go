package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRuntimeListModel(t *testing.T) {
	model := RuntimeList{Total: 5, Runtimes: []Runtime{Runtime{Id: "python-3.8", Key: "python", Name: "Python", Version: "3.8", Base: "python:3.8-alpine", Image: "appwrite\\/runtime-for-python:3.8", Logo: "python.png", Supports: []string{"test"}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result RuntimeList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestRuntimeListModelDefaults(t *testing.T) {
	var fresh RuntimeList
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

	reused := RuntimeList{Total: 5, Runtimes: []Runtime{Runtime{Id: "python-3.8", Key: "python", Name: "Python", Version: "3.8", Base: "python:3.8-alpine", Image: "appwrite\\/runtime-for-python:3.8", Logo: "python.png", Supports: []string{"test"}}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
