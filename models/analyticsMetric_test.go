package models

import (
	"encoding/json"
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
