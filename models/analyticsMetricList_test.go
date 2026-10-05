package models

import (
	"encoding/json"
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
