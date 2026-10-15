package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAnalyticsMetricModel(t *testing.T) {
	model := AnalyticsMetric{Visitors: 1234, Sessions: 4567, Events: 8910}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AnalyticsMetric
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Visitors != model.Visitors {
		t.Errorf("Expected Visitors %v, got %v", model.Visitors, result.Visitors)
	}
	if result.Sessions != model.Sessions {
		t.Errorf("Expected Sessions %v, got %v", model.Sessions, result.Sessions)
	}
	if result.Events != model.Events {
		t.Errorf("Expected Events %v, got %v", model.Events, result.Events)
	}
}

func TestAnalyticsMetricModelDefaults(t *testing.T) {
	var fresh AnalyticsMetric
	if err := json.Unmarshal([]byte("{}"), &fresh); err != nil {
		t.Fatal(err)
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Visitors)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing visitors to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Sessions)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing sessions to default to %v, got %v", want, got)
		}
	}
	{
		var want, got interface{}
		if err := json.Unmarshal([]byte("0"), &want); err != nil {
			t.Fatal(err)
		}
		encoded, err := json.Marshal(fresh.Events)
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(encoded, &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Expected missing events to default to %v, got %v", want, got)
		}
	}

	reused := AnalyticsMetric{Visitors: 1234, Sessions: 4567, Events: 8910}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
