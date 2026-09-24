package models

import (
	"encoding/json"
	"testing"
)

func TestDomainTransferStatusModel(t *testing.T) {
	model := DomainTransferStatus{Status: "pending_registry", Reason: "Transfer in progress", Timestamp: "2020-10-15T06:38:00.000+00:00"}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DomainTransferStatus
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != model.Status {
		t.Errorf("Expected Status %v, got %v", model.Status, result.Status)
	}
	if result.Reason != model.Reason {
		t.Errorf("Expected Reason %v, got %v", model.Reason, result.Reason)
	}
	if result.Timestamp != model.Timestamp {
		t.Errorf("Expected Timestamp %v, got %v", model.Timestamp, result.Timestamp)
	}
}
