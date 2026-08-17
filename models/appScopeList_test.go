package models

import (
	"encoding/json"
	"testing"
)

func TestAppScopeListModel(t *testing.T) {
	model := AppScopeList{Total: 5, Scopes: []AppScope{AppScope{Value: "organization:organization.read", Description: "Access to read the organization", Type: "organization", Category: "Organization", Deprecated: true}}}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result AppScopeList
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != model.Total {
		t.Errorf("Expected Total %v, got %v", model.Total, result.Total)
	}
}
