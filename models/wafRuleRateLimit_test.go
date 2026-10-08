package models

import (
	"encoding/json"
	"testing"
)

func TestWafRuleRateLimitModel(t *testing.T) {
	model := WafRuleRateLimit{Id: "wafRule1", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "Block anonymous POST traffic", Description: "Blocks anonymous POST calls to /v1/graphql", TeamId: "5e5ea5c16897e", ProjectId: "cloudConsole", ResourceType: "functions", ResourceId: "functionId", Action: "deny", Priority: 100, Enabled: true, Conditions: []interface{}{}, Config: map[string]interface{}{}, Limit: 1000, Interval: 60, Key: "userId", Strategy: "slidingWindow", MaxBucketSize: 200}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result WafRuleRateLimit
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
	if result.Description != model.Description {
		t.Errorf("Expected Description %v, got %v", model.Description, result.Description)
	}
	if result.TeamId != model.TeamId {
		t.Errorf("Expected TeamId %v, got %v", model.TeamId, result.TeamId)
	}
	if result.ProjectId != model.ProjectId {
		t.Errorf("Expected ProjectId %v, got %v", model.ProjectId, result.ProjectId)
	}
	if result.ResourceType != model.ResourceType {
		t.Errorf("Expected ResourceType %v, got %v", model.ResourceType, result.ResourceType)
	}
	if result.ResourceId != model.ResourceId {
		t.Errorf("Expected ResourceId %v, got %v", model.ResourceId, result.ResourceId)
	}
	if result.Action != model.Action {
		t.Errorf("Expected Action %v, got %v", model.Action, result.Action)
	}
	if result.Priority != model.Priority {
		t.Errorf("Expected Priority %v, got %v", model.Priority, result.Priority)
	}
	if result.Enabled != model.Enabled {
		t.Errorf("Expected Enabled %v, got %v", model.Enabled, result.Enabled)
	}
	if result.Limit != model.Limit {
		t.Errorf("Expected Limit %v, got %v", model.Limit, result.Limit)
	}
	if result.Interval != model.Interval {
		t.Errorf("Expected Interval %v, got %v", model.Interval, result.Interval)
	}
	if result.Key != model.Key {
		t.Errorf("Expected Key %v, got %v", model.Key, result.Key)
	}
	if result.Strategy != model.Strategy {
		t.Errorf("Expected Strategy %v, got %v", model.Strategy, result.Strategy)
	}
	if result.MaxBucketSize != model.MaxBucketSize {
		t.Errorf("Expected MaxBucketSize %v, got %v", model.MaxBucketSize, result.MaxBucketSize)
	}
}
