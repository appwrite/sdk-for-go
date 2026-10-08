package models

import (
	"encoding/json"
	"testing"
)

func TestWafRuleListModel(t *testing.T) {
	model := WafRuleList{Total: 5, Rules: []WafRule{WafRule{Id: "wafRule1", CreatedAt: "2020-10-15T06:38:00.000+00:00", UpdatedAt: "2020-10-15T06:38:00.000+00:00", Name: "Block anonymous POST traffic", Description: "Blocks anonymous POST calls to /v1/graphql", TeamId: "5e5ea5c16897e", ProjectId: "cloudConsole", ResourceType: "functions", ResourceId: "functionId", Action: "deny", Priority: 100, Enabled: true, Conditions: []interface{}{}, Config: map[string]interface{}{}}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result WafRuleList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
