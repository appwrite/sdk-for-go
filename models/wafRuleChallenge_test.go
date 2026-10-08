package models

import (
	"encoding/json"
	"testing"
)

func TestWafRuleChallengeModel(t *testing.T) {
	model := WafRuleChallenge{Id: "wafRule1", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "Block anonymous POST traffic", Description: "Blocks anonymous POST calls to /v1/graphql", TeamId: "5e5ea5c16897e", ProjectId: "cloudConsole", ResourceType: "functions", ResourceId: "functionId", Action: "deny", Priority: 100, Enabled: true, Conditions: []interface{}{}, Config: map[string]interface{}{}, ChallengeType: "compute", Difficulty: 3, Ttl: 1800}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result WafRuleChallenge
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
	if result.ChallengeType != model.ChallengeType {
		t.Errorf("Expected ChallengeType %v, got %v", model.ChallengeType, result.ChallengeType)
	}
	if result.Difficulty != model.Difficulty {
		t.Errorf("Expected Difficulty %v, got %v", model.Difficulty, result.Difficulty)
	}
	if result.Ttl != model.Ttl {
		t.Errorf("Expected Ttl %v, got %v", model.Ttl, result.Ttl)
	}
}
