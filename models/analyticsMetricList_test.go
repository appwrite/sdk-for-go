package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAnalyticsMetricListModel(t *testing.T) {
	model := AnalyticsMetricList{Total: 30, Metrics: []AnalyticsMetric{AnalyticsMetric{Visitors: 1234, Sessions: 4567, Events: 8910}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AnalyticsMetricList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}

func TestAnalyticsMetricListModelDefaults(t *testing.T) {
	var fresh AnalyticsMetricList
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

	reused := AnalyticsMetricList{Total: 30, Metrics: []AnalyticsMetric{AnalyticsMetric{Visitors: 1234, Sessions: 4567, Events: 8910}}}
	if err := json.Unmarshal([]byte("{}"), &reused); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(fresh, reused) {
		t.Errorf("Expected decoding into a used model to match a fresh decode, got %+v", reused)
	}
}
