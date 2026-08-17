package models

import (
	"encoding/json"
	"testing"
)

func TestDatabaseStatusConnectionsModel(t *testing.T) {
	model := DatabaseStatusConnections{Current: 12, Max: 100}

	data, err := json.Marshal(model)
	if err != nil {
		t.Fatal(err)
	}

	var result DatabaseStatusConnections
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal(err)
	}
	if result.Current != model.Current {
		t.Errorf("Expected Current %v, got %v", model.Current, result.Current)
	}
	if result.Max != model.Max {
		t.Errorf("Expected Max %v, got %v", model.Max, result.Max)
	}
}
