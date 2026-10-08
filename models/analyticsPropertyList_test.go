package models

import (
	"encoding/json"
	"testing"
)

func TestAnalyticsPropertyListModel(t *testing.T) {
	model := AnalyticsPropertyList{Total: 5, Properties: []AnalyticsProperty{AnalyticsProperty{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "My Website", Domain: "example.com", Enabled: true, Public: true, AllowedOrigins: []string{"test"}, AccessedAt: "2020-10-15T06:38:00.000+00:00", FirstAccessedAt: "2020-10-15T06:38:00.000+00:00"}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AnalyticsPropertyList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
