package models

import (
	"encoding/json"
	"testing"
)

func TestWafRuleListModel(t *testing.T) {
	model := WafRuleList{Total: 5, Rules: []interface{}{}}

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
