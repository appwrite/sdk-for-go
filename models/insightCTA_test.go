package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestInsightCTAModel(t *testing.T) {
	model := InsightCTA{Label: "Create missing index", Service: "tablesDB", Method: "createIndex", Params: map[string]interface{}{}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result InsightCTA
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Label != model.Label {
		t.Errorf("Expected Label %v, got %v", model.Label, result.Label)
	}
	if result.Service != model.Service {
		t.Errorf("Expected Service %v, got %v", model.Service, result.Service)
	}
	if result.Method != model.Method {
		t.Errorf("Expected Method %v, got %v", model.Method, result.Method)
	}
}

func TestInsightCTAModelDefaults(t *testing.T) {
	var fresh InsightCTA
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Label)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing label to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Service)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing service to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("\"\""), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Method)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing method to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("[]"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Params)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing params to default to %v, got %v", want, got)
		}
	}

	reused := InsightCTA{Label: "Create missing index", Service: "tablesDB", Method: "createIndex", Params: map[string]interface{}{}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
