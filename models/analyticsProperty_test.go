package models

import (
	"encoding/json"
	"testing"
)

func TestAnalyticsPropertyModel(t *testing.T) {
	model := AnalyticsProperty{Id: "5e5ea5c16897e", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "My Website", Domain: "example.com", Enabled: true, Public: true, AllowedOrigins: []string{"test"}, AccessedAt: "2020-10-15T06:38:00.000+00:00", FirstAccessedAt: "2020-10-15T06:38:00.000+00:00"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AnalyticsProperty
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != model.Id {
		t.Errorf("Expected Id %v, got %v", model.Id, result.Id)
	}
	if result.CreatedAt != model.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", model.CreatedAt, result.CreatedAt)
	}
	if result.UpdatedAt != model.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", model.UpdatedAt, result.UpdatedAt)
	}
	if result.Name != model.Name {
		t.Errorf("Expected Name %v, got %v", model.Name, result.Name)
	}
	if result.Domain != model.Domain {
		t.Errorf("Expected Domain %v, got %v", model.Domain, result.Domain)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
	if result.Public != model.Public {
		t.Errorf("Expected Public %v, got %v", model.Public, result.Public)
	}
	if result.AccessedAt != model.AccessedAt {
		t.Errorf("Expected AccessedAt %v, got %v", model.AccessedAt, result.AccessedAt)
	}
	if result.FirstAccessedAt != model.FirstAccessedAt {
		t.Errorf("Expected FirstAccessedAt %v, got %v", model.FirstAccessedAt, result.FirstAccessedAt)
	}
}
