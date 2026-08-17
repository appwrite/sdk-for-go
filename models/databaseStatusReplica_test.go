package models

import (
	"encoding/json"
	"testing"
)

func TestDatabaseStatusReplicaModel(t *testing.T) {
	model := DatabaseStatusReplica{Index: 0, Role: "primary", Healthy: true}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DatabaseStatusReplica
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Index != model.Index {
		t.Errorf("Expected Index %v, got %v", model.Index, result.Index)
	}
	if result.Role != model.Role {
		t.Errorf("Expected Role %v, got %v", model.Role, result.Role)
	}
	if result.Healthy != model.Healthy {
		t.Errorf("Expected Healthy %v, got %v", model.Healthy, result.Healthy)
	}
}
